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
    {{{1, 1, 1, 1}}, 0, 0},         // I
    {{{1, 1, 1}, {0, 1, 0}}, 0, 0}, // T
    {{{1, 1, 0}, {0, 1, 1}}, 0, 0}, // Z
    {{{0, 1, 1}, {1, 1, 0}}, 0, 0}, // S
    {{{1, 1}, {1, 1}}, 0, 0},       // O
};

int currentTetrominoIndex = 0;

void DrawBoard() {
  for (int y = 0; y < BOARD_HEIGHT; y++) {
    for (int x = 0; x < BOARD_WIDTH; x++) {
      if (board[y][x]) {
        DrawRectangle(x * BLOCK_SIZE, y * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE,
                      BLUE);
      }
    }
  }
}

void DrawTetromino(Tetrimino t) {
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      if (t.shape[y][x]) {
        DrawRectangle((t.x + x) * BLOCK_SIZE, (t.y + y) * BLOCK_SIZE,
                      BLOCK_SIZE, BLOCK_SIZE, RED);
      }
    }
  }
}

void RotateTetromino(Tetrimino *t) {
  int temp[4][4] = {0};
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      temp[x][3 - y] = t->shape[y][x];
    }
  }
  // Check if the rotation is valid
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      if (temp[y][x] &&
          (t->x + x < 0 || t->x + x >= BOARD_WIDTH ||
           t->y + y >= BOARD_HEIGHT || board[t->y + y][t->x + x])) {
        return; // Invalid rotation
      }
    }
  }
  // Apply rotation
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      t->shape[y][x] = temp[y][x];
    }
  }
}

bool CheckCollision(Tetrimino t) {
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      if (t.shape[y][x] &&
          (t.x + x < 0 || t.x + x >= BOARD_WIDTH || t.y + y >= BOARD_HEIGHT ||
           board[t.y + y][t.x + x])) {
        return true; // Collision detected
      }
    }
  }
  return false;
}

void LockTetromino(Tetrimino t) {
  for (int y = 0; y < 4; y++) {
    for (int x = 0; x < 4; x++) {
      if (t.shape[y][x]) {
        board[t.y + y][t.x + x] = 1; // Lock the tetromino in place
      }
    }
  }
}

void ClearLines() {
  for (int y = BOARD_HEIGHT - 1; y >= 0; y--) {
    bool fullLine = true;
    for (int x = 0; x < BOARD_WIDTH; x++) {
      if (!board[y][x]) {
        fullLine = false;
        break;
      }
    }
    if (fullLine) {
      // Clear the line
      for (int row = y; row > 0; row--) {
        for (int x = 0; x < BOARD_WIDTH; x++) {
          board[row][x] = board[row - 1][x];
        }
      }
      // Clear the top row
      for (int x = 0; x < BOARD_WIDTH; x++) {
        board[0][x] = 0;
      }
      y++; // Check this row again
    }
  }
}

void UpdateGame() {
  static float timer = 0;
  static float delay = 1.0f; // Delay for falling speed
  timer += GetFrameTime();

  Tetrimino currentTetromino = tetrominos[currentTetrominoIndex];
  if (timer >= delay) {
    timer = 0;
    // Move tetromino down

    currentTetromino.y += 1;
    if (CheckCollision(currentTetromino)) {
      currentTetromino.y -= 1; // Revert position
      LockTetromino(currentTetromino);
      ClearLines();
      currentTetrominoIndex =
          GetRandomValue(0, sizeof(tetrominos) / sizeof(tetrominos[0]) - 1);
      currentTetromino.x = BOARD_WIDTH / 2 - 2; // Reset position
      currentTetromino.y = 0;

      // Check for game over condition
      if (CheckCollision(currentTetromino)) {
        // Game over logic
      }
    }
  }

  // Handle input for left, right, and rotate
  if (IsKeyPressed(KEY_LEFT)) {
    currentTetromino.x -= 1;
    if (CheckCollision(currentTetromino))
      currentTetromino.x += 1; // Revert
  }
  if (IsKeyPressed(KEY_RIGHT)) {
    currentTetromino.x += 1;
    if (CheckCollision(currentTetromino))
      currentTetromino.x -= 1; // Revert
  }
  if (IsKeyPressed(KEY_DOWN)) {
    currentTetromino.y += 1;
    if (CheckCollision(currentTetromino))
      currentTetromino.y -= 1; // Revert
  }
  if (IsKeyPressed(KEY_UP)) {
    RotateTetromino(&currentTetromino);
  }
}

int main() {
  InitWindow(BOARD_WIDTH * BLOCK_SIZE, BOARD_HEIGHT * BLOCK_SIZE,
             "Tetris in Raylib");
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
