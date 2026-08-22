The Meridian Pivot
Assignment 1 — Mini-Prototype & Learning/Blocker Journal
Group G29
Sprint 2 · Tool: Go (Golang)
1. Tool Assigned
Go (Golang) — selected specifically due to lack of prior familiarity. Learning focused on foundational syntax (structs, slices, loops, functions, and conditionals) before applying these concepts to build an HTTP webhook-based synchronization prototype with HMAC-SHA256 signature verification.
2. Time-box vs Actual
Time-boxed: 6 hours (Day 1–2 combined) for basic language acquisition and a working, self-tested prototype.
Actual: Approximately 5 hours 50 minutes across the logged entries below — within the time-box.
3. Prototype Summary
Deliverable: A Go HTTP server using net/http and crypto/hmac that:
Listens for inbound webhook payload notifications.
Recomputes the expected HMAC-SHA256 signature over the raw request payload.
Rejects unauthorized requests missing a valid signature header with HTTP 401.
Accepts verified requests and returns HTTP 200.
Companion Setup: A test sender script firing valid and invalid webhook requests against the Go server to verify authentication outcomes.
Run Instructions: Run go run main.go in one terminal, then trigger test requests via the sender script/cURL in a second terminal.
4. Blocker Journal
Logged in real time as each issue was hit, not reconstructed afterward.
Time
What happened
Diagnosis
Fix / Resolution
Time lost
Day 1, 10:00
Hit persistent Go syntax and type errors during basic setup.
Tried to write code using conventions from familiar dynamic/object-oriented languages rather than idiomatic Go patterns.
Isolated syntax logic by writing small, standalone Go example scripts for structs, slices, and functions before re-integrating.
1 hr 30 min
Day 1, 14:15
Server continuously rejected valid webhooks with an HTTP 401 Unauthorized status.
Header mismatch: the sender script sent X-Webhook-Signature, but the Go net/http handler expected X-Signature.
Standardized header names and signature generation logic to match on both client and server sides.
1 hr 30 min
Day 2, 09:30
Server startup crashed with listen tcp :8080: bind: Only one usage of each socket address.
A zombie process or background service was already bound to port 8080.
Identified the process using Get-NetTCPConnection in PowerShell and killed the conflicting process ID.
30 min
Day 2, 11:00
Go HTTP handler panicked/logged errors regarding superfluous response writes.
Executed multiple http.Error() or w.WriteHeader() calls in a single execution path without returning.
Added explicit return statements immediately following each response write call to halt handler execution.
20 min
Day 2, 12:00
Ran full verification test suite across valid and invalid request cases.
n/a — confirmation step.
All test cases returned expected status codes (200 OK / 401 Unauthorized). Prototype complete.
20 min

5. Verification Run
Executed test sender suite against the running Go HTTP server. Results:
Valid HMAC Signature & Matching Header: HTTP 200 OK
Invalid/Mismatched HMAC Signature: HTTP 401 Unauthorized
Missing Header Name Alignment: HTTP 401 Unauthorized
All outcomes matched expectations on the first full run after resolving execution control flow and header alignment.


6. Reflection
Building in an unfamiliar compiled language reinforced that debugging requires understanding the underlying error and language behavior rather than blindly tweaking code. My single biggest takeaway was the importance of isolating request flow design early — attempting to debug HTTP routing, HMAC cryptography, and language syntax simultaneously cost unnecessary time. Writing isolated unit tests for smaller Go components prior to integration will be my primary procedural shift in future sprints.

