tetris: tetris.c
	cc `pkg-config --cflags raylib` `pkg-config --libs raylib` tetris.c -o tetris
