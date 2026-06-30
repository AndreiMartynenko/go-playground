

Week 2 — Request Processing & Concurrency

Why do we need concurrency?

Modern applications handle thousands of requests simultaneously.

Without concurrency, each request would have to wait until the previous one finishes.

Example:

User A -----> Server -----> Database (waiting...)
User B
waits...
User C
waits...

This wastes CPU time because the CPU spends most of its time 
waiting for external resources.

⸻

CPU vs Memory

CPU speed has increased much faster than RAM speed.

Because of this, modern CPUs use caches.

Fastest
L1 Cache
L2 Cache
L3 Cache
RAM
Slowest

Cache is much smaller but much faster than RAM.

⸻

Context Switching

A CPU core executes only one instruction at a time.

When switching from one thread to another, the operating system must:

* save current thread state
* restore another thread state
* reload CPU registers
* sometimes reload memory cache

This operation is called a Context Switch.

Context switching is expensive.

Too many OS threads = worse performance.

⸻

Evolution of Web Servers

1. CGI (Process per Request)

Request
↓
Create Process
↓
Handle Request
↓
Destroy Process

Advantages:

* Simple

Disadvantages:

* Very expensive
* High memory usage
* Slow

⸻

2. Multithreading

Request 1 → Thread 1
Request 2 → Thread 2
Request 3 → Thread 3

Advantages:

* Faster than CGI

Disadvantages:

* Threads are expensive
* Context switching
* High memory usage

⸻

3. Event Loop (Node.js)

One thread handles many requests using non-blocking I/O.

Thread
↓
Request A (waiting DB)
↓
Request B
↓
Request C

Advantages:

* Very low memory usage
* Excellent for I/O

Disadvantages:

* CPU-heavy tasks block everything.

⸻

4. Go (Goroutines)

Go combines the best ideas.

Thousands of Goroutines
↓
Go Scheduler
↓
Small number of OS Threads
↓
CPU

Advantages:

* lightweight
* cheap
* automatic scheduling
* excellent scalability

⸻

CPU Work vs Waiting

A request usually looks like this:

CPU
↓
Database
↓
CPU
↓
API
↓
CPU

Graphically:

CPU   WAIT   CPU    WAIT    CPU
████░░░░░████░░░░░░░████

Notice:

Most of the request is actually waiting.

⸻

Why Goroutines Exist

Instead of wasting CPU while one request waits:

Request A
Waiting for DB
↓
CPU is idle

Go allows another goroutine to use the CPU.

Goroutine A
Waiting
↓
Scheduler
↓
Run Goroutine B
↓
Run Goroutine C

The CPU stays busy.

⸻

Blocking vs Non-Blocking

Blocking

CPU
↓
Wait
↓
Continue

The thread cannot do anything while waiting.

⸻

Non-blocking

CPU
↓
Start operation
↓
Continue doing other work
↓
Receive result later

This improves CPU utilization.

⸻

What is Concurrency?

Concurrency means:

Multiple tasks make progress during the same period of time.

They are not necessarily running simultaneously.

⸻

What is Parallelism?

Parallelism means:

Multiple tasks execute at the exact same time.

Requires multiple CPU cores.

Example:

Core 1 → Task A
Core 2 → Task B
Core 3 → Task C

⸻

Concurrency vs Parallelism

Concurrency

One CPU
Task A
Task B
Task C
Scheduler switches between them.

Parallelism

CPU1 → Task A
CPU2 → Task B
CPU3 → Task C

⸻

Why Go is Fast

Go uses:

* Goroutines (lightweight threads)
* Scheduler (maps goroutines to OS threads)
* Channels (safe communication)
* Runtime (manages execution automatically)

Developers write concurrent code without managing OS threads manually.

⸻

Key Takeaways

✅ CPU is much faster than RAM.

✅ Context switching between OS threads is expensive.

✅ Most backend requests spend most of their lifetime waiting.

✅ Threads are heavy.

✅ Goroutines are lightweight.

✅ Go Scheduler automatically manages goroutines.

✅ Concurrency improves CPU utilization.

✅ Parallelism requires multiple CPU cores.

⸻

Interview Summary (Very Important)

If an interviewer asks:

Why are goroutines useful?

A strong answer is:

Goroutines are lightweight concurrent execution units managed by the Go runtime. 
Unlike OS threads, they require very little memory and are scheduled efficiently by the Go scheduler. 
Since backend applications spend most of their time waiting for 
I/O operations such as databases, APIs, or network calls, goroutines allow the CPU 
to execute other work instead of remaining idle. This enables Go applications 
to handle thousands of concurrent requests with much lower overhead than 
traditional thread-based models.

⸻

