package com.example;

public class Hello {
    // Load the shared library (libhello.dylib or libhello.so)
    static {
        System.loadLibrary("hello");  
        // if you named it libdispatch.dylib, use:
        // System.loadLibrary("dispatch");
    }

    // Declaration must match: Java_com_example_Hello_SayHello(JNIEnv*, jclass, jstring)
    private native String SayHello(String input);

    public static void main(String[] args) throws InterruptedException {
        Hello h = new Hello();
        int numThreads = 100;
        Thread[] threads = new Thread[numThreads];
        for (int i = 0; i < numThreads; i++) {
            threads[i] = new Thread(() -> {
                String result = h.SayHello("thread-" + Thread.currentThread().getId());
                System.out.println(Thread.currentThread().getName() + ": " + result);
            });
            threads[i].start();
        }
        for (Thread t : threads) {
            t.join();
        }
    }
}