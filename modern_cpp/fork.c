#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
int main(void)
{
   int i;
   for(i=0; i<2; i++){
      fork();
      /* write(1, "-", 1); */
      printf("printf-\n");
      printf("print-");
   }
   wait(NULL);
   wait(NULL);
   return 0;
}
