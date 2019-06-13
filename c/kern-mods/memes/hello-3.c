/*
 * hello-3.c - __init, __initdata __exit macros
 */
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>

static int hello3_data __initdata = 3;

static int __init hello_3_init(void) {
  printk(KERN_INFO "Hello, world %d\n", hello3_data);
  return 0;
}

static void __exit hello_3_exit(void) { printk(KERN_INFO "Goodbye, world\n"); }

module_init(hello_3_init);
module_exit(hello_3_exit);
