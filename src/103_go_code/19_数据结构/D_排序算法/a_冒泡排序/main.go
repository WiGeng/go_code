package main

/* 冒泡排序
冒泡排序 （Bubble Sorting）的基本思想是：通过对待排序序列从后向前（从下标较大的元素开始），依次比较相邻元素的排序码，若发现逆序
则交换，使排序码较小的元素逐渐从后部移向前部（从下标较大的单元移向下标较小的单元），就象水底下的气泡一样逐渐向上冒。

因为排序的过程中，各元素不断接近自己的位置，如果一趟比较下来没有进行过交换，就说明序列有序，因此要在排序过程中设置一个标志
flag判断元素是否进行过交换,从而减少不必要的比较。
*/
