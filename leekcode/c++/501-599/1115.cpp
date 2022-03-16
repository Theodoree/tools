class FooBar {
private:
    int n;
    std::atomic_int64_t m_flag;

public:
    FooBar(int n) {
        m_flag = 0;
        this->n = n;
    }

    void foo(function<void()> printFoo) {

        for (int i = 0; i < n; i++) {
            while (m_flag != 0){
                 this_thread::yield();
            };
        	// printFoo() outputs "foo". Do not change or remove this line.
        	printFoo();
            m_flag = 1;
        }
    }

    void bar(function<void()> printBar) {

        for (int i = 0; i < n; i++) {
               while (m_flag != 1){
                   this_thread::yield();
               };
        	// printBar() outputs "bar". Do not change or remove this line.
        	printBar();
            m_flag = 0;
        }
    }
};