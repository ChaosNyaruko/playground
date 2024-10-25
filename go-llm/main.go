// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command ragserver is an HTTP server that implements RAG (Retrieval
// Augmented Generation) using the Gemini model and Weaviate. See the
// accompanying README file for additional details.
package main

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

const generativeModelName = "llama3"
const embeddingModelName = "llama3"

type ollama struct {
	model string
}

func (m *ollama) Generate(ctx context.Context, text string) (string, error) {
	data := map[string]any{"model": m.model, "prompt": text, "stream": false}
	js, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("marshal err in generate: %v", err)
	}

	/*
		{
		  "model": "llama3.2",
		  "created_at": "2023-08-04T19:22:45.499127Z",
		  "response": "The sky is blue because it is the color of the sky.",
		  "done": true,
		  "context": [1, 2, 3],
		  "total_duration": 5043500667,
		  "load_duration": 5025959,
		  "prompt_eval_count": 26,
		  "prompt_eval_duration": 325953000,
		  "eval_count": 290,
		  "eval_duration": 4709213000
		}
	*/
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(js))
	if err != nil {
		return "", fmt.Errorf("query ollama generate err: %v", err)
	}
	x, err := io.ReadAll(resp.Body)
	res := map[string]any{}
	if err := json.Unmarshal(x, &res); err != nil {
		return "", fmt.Errorf("unmarshal error in generate %v", err)
	}
	return res["response"].(string), nil
}

func (m *ollama) Embed(ctx context.Context, docs []string) ([][]float32, error) {
	data := map[string]any{"model": m.model, "input": docs, "stream": false}
	js, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal err in generate: %v", err)
	}

	if err != nil {
	}
	/*
		curl http://localhost:11434/api/embed -d '{
		  "model": "all-minilm",
		  "input": ["Why is the sky blue?", "Why is the grass green?"]
		}'
	*/
	resp, err := http.Post("http://localhost:11434/api/embed", "application/json", bytes.NewReader(js))
	if err != nil {
		return nil, fmt.Errorf("query ollama generate err: %v", err)
	}
	x, err := io.ReadAll(resp.Body)
	type EmbeddingResp struct {
		Model      string
		Embeddings [][]float32
	}
	/*
			{
		  "model": "all-minilm",
		  "embeddings": [[
		    0.010071029, -0.0017594862, 0.05007221, 0.04692972, 0.054916814,
		    0.008599704, 0.105441414, -0.025878139, 0.12958129, 0.031952348
		  ],[
		    -0.0098027075, 0.06042469, 0.025257962, -0.006364387, 0.07272725,
		    0.017194884, 0.09032035, -0.051705178, 0.09951512, 0.09072481
		  ]]
		}
	*/
	var res EmbeddingResp
	if err := json.Unmarshal(x, &res); err != nil {
		return nil, fmt.Errorf("unmarshal error in generate %v", err)
	}
	return res.Embeddings, nil
}

// This is a standard Go HTTP server. Server state is in the ragServer struct.
// The `main` function connects to the required services (Weaviate and Google
// AI), initializes the server state and registers HTTP handlers.
func main() {
	ctx := context.Background()
	wvClient, err := initWeaviate(ctx)
	if err != nil {
		log.Fatal(err)
	}

	llama3 := &ollama{model: generativeModelName}

	server := &ragServer{
		ctx:      ctx,
		wvClient: wvClient,
		genModel: llama3,
		embModel: llama3,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type GenLLM interface {
	Generate(context.Context, string) (string, error)
}

type EmbeddingModel interface {
	Embed(context.Context, []string) ([][]float32, error)
}

type ragServer struct {
	ctx      context.Context
	wvClient *weaviate.Client
	genModel GenLLM
	embModel EmbeddingModel
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type document struct {
		Text string
	}
	type addRequest struct {
		Documents []document
	}
	ar := &addRequest{}

	err := readRequestJSON(req, ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the batch embedding API to embed all documents at once.
	batch := []string{}
	for _, doc := range ar.Documents {
		batch = append(batch, doc.Text)
	}
	log.Printf("invoking embedding model with %v documents", len(ar.Documents))
	rsp, err := rs.embModel.Embed(rs.ctx, batch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(rsp) != len(ar.Documents) {
		http.Error(w, "embedded batch size mismatch", http.StatusInternalServerError)
		return
	}

	// Convert our documents - along with their embedding vectors - into types
	// used by the Weaviate client library.
	objects := make([]*models.Object, len(ar.Documents))
	for i, doc := range ar.Documents {
		log.Printf("object: %d, embeded len: %d", i, len(rsp[i]))
		objects[i] = &models.Object{
			Class: "Document",
			Properties: map[string]any{
				"text": doc.Text,
			},
			Vector: rsp[i],
		}
	}

	// Store documents with embeddings in the Weaviate DB.
	log.Printf("storing %v objects in weaviate", len(objects))
	_, err = rs.wvClient.Batch().ObjectsBatcher().WithObjects(objects...).Do(rs.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type queryRequest struct {
		Content string
	}
	qr := &queryRequest{}
	err := readRequestJSON(req, qr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Embed the query contents.
	rsp, err := rs.embModel.Embed(rs.ctx, []string{qr.Content})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search weaviate to find the most relevant (closest in vector space)
	// documents to the query.
	gql := rs.wvClient.GraphQL()
	result, err := gql.Get().
		WithNearVector(
			gql.NearVectorArgBuilder().WithVector(rsp[0])).
		WithClassName("Document").
		WithFields(graphql.Field{Name: "text"}).
		WithLimit(3).
		Do(rs.ctx)
	if werr := combinedWeaviateError(result, err); werr != nil {
		http.Error(w, werr.Error(), http.StatusInternalServerError)
		return
	}

	contents, err := decodeGetResults(result)
	if err != nil {
		http.Error(w, fmt.Errorf("reading weaviate response: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	// Create a RAG query for the LLM with the most relevant documents as
	// context.
	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(contents, "\n"))
	resp, err := rs.genModel.Generate(rs.ctx, ragQuery)
	if err != nil {
		log.Printf("calling generative model: %v", err.Error())
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	var respTexts []string = []string{resp}
	// for _, part := range resp.Candidates[0].Content.Parts {
	// 	if pt, ok := part.(genai.Text); ok {
	// 		respTexts = append(respTexts, string(pt))
	// 	} else {
	// 		log.Printf("bad type of part: %v", pt)
	// 		http.Error(w, "generative model error", http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	renderJSON(w, strings.Join(respTexts, "\n"))
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`

// decodeGetResults decodes the result returned by Weaviate's GraphQL Get
// query; these are returned as a nested map[string]any (just like JSON
// unmarshaled into a map[string]any). We have to extract all document contents
// as a list of strings.
func decodeGetResults(result *models.GraphQLResponse) ([]string, error) {
	data, ok := result.Data["Get"]
	if !ok {
		return nil, fmt.Errorf("Get key not found in result")
	}
	doc, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("Get key unexpected type")
	}
	slc, ok := doc["Document"].([]any)
	if !ok {
		return nil, fmt.Errorf("Document is not a list of results")
	}

	var out []string
	for _, s := range slc {
		smap, ok := s.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid element in list of documents")
		}
		s, ok := smap["text"].(string)
		if !ok {
			return nil, fmt.Errorf("expected string in list of documents")
		}
		out = append(out, s)
	}
	return out, nil
}
