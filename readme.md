Project Goal & Scenarios: Health Check Resilience and Concurrency

Purpose

The primary objective of this project is to develop a highly performant and resilient health check mechanism that accurately simulates real-world dependencies. The implementation will focus on benchmarking various concurrency strategies to evaluate their impact on overall latency and failure handling under realistic load conditions.

Disclaimer

This codebase was developed without the aid of generative AI tools. This ensures the solution is a genuine demonstration of technical proficiency, problem-solving skills, and hands-on experience in managing concurrency and I/O bottlenecks.

Performance Scenarios

The project is designed to simulate and measure the performance characteristics under three distinct concurrency models:

Sequential Execution (No Parallelism):

Description: All dependency health checks are executed one after the other.

Goal: To measure the performance cost of synchronous calls, where the total execution time is the cumulative sum of all individual dependency latencies ($\sum t_i$).

Unbounded Parallelism (No Timeout):

Description: All dependency health checks are initiated concurrently (pipelining is maximized).

Goal: To measure the worst-case scenario where the total execution time is dictated by the single slowest dependency, resulting in high latency variability. The total time is $\max(t_1, t_2, \ldots, t_n)$.

Bounded Parallelism with Latency Budget (Timeout):

Description: All dependency health checks are initiated concurrently, but a fixed, aggressive timeout (latency budget) is imposed on the entire operation.

Goal: To showcase resilience and partial failure handling. The system will stop waiting once the budget is exhausted, and the measured result will be the successful responses received within the allotted time, prioritizing user experience over full report fidelity.