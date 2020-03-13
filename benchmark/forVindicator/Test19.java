package test;

public class Test19 extends Thread {
  static int x;
  static int a;
  static int y;
  static int z;
  static final Object m = new Object();
  static final Object n = new Object();

  static void sleepSec(float sec) {
    try {
	Thread.sleep((long)(sec*1000));
    } catch (Exception e) {
       throw new RuntimeException(e);
    }
  }


  @Override
  public void run() {
    //T1
	z = 1;
	System.out.println("w(z)");
    synchronized(m) {
        System.out.println("l(m)");
        x = 1;
        System.out.println("w(x)");
        y = 2;
        System.out.println("w(y)");
        System.out.println("u(m)");
    }
  } 

   public static class T2 extends Thread implements Runnable {
		public void run() {
			sleepSec(5);
			synchronized(m) {
                System.out.println("l(m)");
                x = 1;
                System.out.println("w(x)");
                System.out.println("u(m)");
			}
            z = 2;
            System.out.println("w(z)");
		}
   }

   public static class T3 extends Thread implements Runnable {
		public void run() {
			sleepSec(10);
			int tmp = x;
            System.out.println("r(x)");
            int tmp2 = y;
            System.out.println("r(y)");
		}
   }

   public static void main(String args[]) throws Exception {
	final Test19 thread1 = new Test19();
    final T2 thread2 = new T2();
    final T3 thread3 = new T3();

	thread1.start();
	thread2.start();
	thread3.start();

	thread1.join();
	thread2.join();
	thread3.join();

   }
}
