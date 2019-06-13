#include <ncurses.h>

int screenSetUp();
int mapSetUp();

int screenSetUp() {
  initscr();
  printw("Hello world!");
  noecho();
  refresh();

  return 0;
}

int mapSetUp() {
  mvprintw()

      return 0;
}

int main() {
  screenSetUp();
  getch();
  endwin();
  return 0;
}
