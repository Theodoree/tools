#include <semaphore.h>

class ZeroEvenOdd {
private:
    int n;
    sem_t printOdd, printEven, numDone;
public:
    ZeroEvenOdd(int n) {
        this->n = n;
        sem_init(&printOdd, 0, 0);
        sem_init(&printEven, 0, 0);
        sem_init(&numDone, 0, 1);
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber) {
        for(int i = 0; i < n; i++){
            sem_wait(&numDone);
            printNumber(0);
            if(i % 2 == 0){
                sem_post(&printOdd);
            }else{
                sem_post(&printEven);
            }
        }
    }

    void even(function<void(int)> printNumber) {
        for (int i = 2; i <= n; i += 2) {
            sem_wait(&printEven);
            printNumber(i);
            sem_post(&numDone);
        }
    }

    void odd(function<void(int)> printNumber) {
        for (int i = 1; i <=n; i += 2) {
            sem_wait(&printOdd);
            printNumber(i);
            sem_post(&numDone);
        }
    }
};
