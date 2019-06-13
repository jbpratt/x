#include <stdio.h>
#include <stdlib.h>

char *read_line(void) {
  char *line = NULL;
  ssize_t buffsize = 0;
  getline(&line, &buffsize, stdin);
  return line;
}

void loop(void) {
  int status;
  char *line;
  char **args;

  do {
    printf("$ ");
    line = read_line();
    printf(line);

    free(line);
    free(args);
  } while (status);
}

int main(int argc, char **argv) {
  // search for config
  // load it

  // command loop
  loop();
  // cleanup

  return EXIT_SUCCESS;
}
