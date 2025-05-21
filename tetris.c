#include "raylib.h"

#define BOARD_WIDTH 10
#define BOARD_HEIGHT 20
#define BLOCK_SIZE 30

int board[BOARD_HEIGHT][BOARD_WIDTH] = {0};

typedef struct {
    int shape[4][4];
    int x, y;
} Tetrimino;

// Define Tetromino shapes
Tetrimino tetrominos[] = {
    {{{1, 1, 1, 1}}, 0, 0}, // I
    {{{1, 1, 1}, {0, 1, 0}}, 0, 0}, // T
    {{{1, 1, 0}, {0, 1, 1}}, 0, 0}, // Z
    {{{0, 1, 1}, {1, 1, 0}}, 0, 0}, // S
    {{{1, 1}, {1, 1}}, 0, 0}, // O
};

int currentTetrominoIndex = 0;

void DrawBoard() {
    for (int y = 0; y < BOARD_HEIGHT; y++) {
        for (int x = 0; x < BOARD_WIDTH; x++) {
            if (board[y][x]) {
                DrawRectangle(x * BLOCK_SIZE, y * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE, BLUE);
            }
        }
    }
}

void DrawTetromino(Tetrimino t) {
    for (int y = 0; y < 4; y++) {
        for (int x = 0; x < 4; x++) {
            if (t.shape[y][x]) {
                DrawRectangle((t.x + x) * BLOCK_SIZE, (t.y + y) * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE, RED);
            }
        }
    }
}

void UpdateGame() {
    // Move down the tetromino
    // Handle input and rotation
    // Check for collision and line clearing
}

int main() {
    InitWindow(BOARD_WIDTH * BLOCK_SIZE, BOARD_HEIGHT * BLOCK_SIZE, "Tetris in Raylib");
    SetTargetFPS(60);

    while (!WindowShouldClose()) {
        UpdateGame();

        BeginDrawing();
        ClearBackground(RAYWHITE);
        DrawBoard();
        DrawTetromino(tetrominos[currentTetrominoIndex]);
        EndDrawing();
    }

    CloseWindow();
    return 0;
}
