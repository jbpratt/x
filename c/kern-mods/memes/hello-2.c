/*
 *  hello-2.c - Use module_init() and module_exit() macros
 *  rather than using init_module() and cleanup_module()
 */
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>

static int __init hello_2_init(void) {
  printk(KERN_INFO "Hello, world 2\n");
  return 0;
}

static void __exit hello_2_exit(void) {
  printk(KERN_INFO "Goodbye, world 2\n");
}

module_init(hello_2_init);
module_exit(hello_2_exit);
