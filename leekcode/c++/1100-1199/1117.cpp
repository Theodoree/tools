

#include <semaphore.h>
class Sem {
private:
    int n_;
    mutex mu_;
    condition_variable cv_;

public:
    Sem(int n = 0): n_{n} {}

public:
    void wait(int t = 1) {
        unique_lock<mutex> lock(mu_);
        if (n_ < t) {
            cv_.wait(lock, [this, t]{return n_ >= t;});
        }
        n_ -= t;
    }

    void post(int t = 1) {
        unique_lock<mutex> lock(mu_);
        n_ += t;
        if(t > 1)cv_.notify_all();
        else cv_.notify_one();
    }
};

class H2O {
public:
    Sem h_need, h_now, o_need, o_now;
    H2O() : h_need(2), o_need(1){ }

    void hydrogen(function<void()> releaseHydrogen) {
        h_need.wait();
        o_now.post();
        h_now.wait();
        releaseHydrogen();
        h_need.post();
    }

    void oxygen(function<void()> releaseOxygen) {
        o_need.wait();
        o_now.wait(2);
        h_now.post(2);
        releaseOxygen();
        o_need.post();
    }
};