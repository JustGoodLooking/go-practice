    main()
      │
      ├──▶ producer() ─────┐
      │                   ▼
      ├──▶ worker 1 ─────────▶ jobs chan ▶──┐
      ├──▶ worker 2 ─────────▶             ▼
      ├──▶ worker 3 ─────────▶             results chan
      │
[Ctrl+C] → cancel() → producer 停止 → 關 jobs chan
                         │
                    workers 持續做完剩下的 job
                         │
                     每個 worker 做完 → wg.Done()
                         │
                    main() wait 完成 → close(results)
                         │
                       結束 🎉
