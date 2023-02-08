package main

/* 反射
基本介绍
	1) 反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)、类别(kind)
	2) 如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法）
	3) 通过反射，可以修改变量的值，可以调用关联的方法
	4) 使用反射，需要 import("reflect")

引用包 package reflect
import "reflect"
	reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取
	其动态类型信息，该函数返回一个Type类型值。调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。Zero接受一个Type
	类型参数并返回一个代表该类型零值的Value类型值。

应用场景
反射常见应用场景有以下两种
	1）不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法反射。例如以下这种桥接模式，比如我
	前面提出问题。
	""" func bridge(functr interface{}, args ...interface(}) """
	第一个参数funcPtr以接口的形式传入函数指针，函数参数args以可变参数的形式传入，bridge函数中可以用反射来动态执行funcPtr函数
	2）对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串

反射重要的函数和概念
	1) reflect.TypeOf(变量名），获取变量的类型，返回reflect.Type类型
	2) reflect.ValueOf(变量名），获取变量的值，返回reflect.Value类型reflect.value是一个结构体类型。通过reflect.Value，可以获取到关于该变量的很多信息。
*/
