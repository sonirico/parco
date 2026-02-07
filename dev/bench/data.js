window.BENCHMARK_DATA = {
  "lastUpdate": 1770486302525,
  "repoUrl": "https://github.com/sonirico/parco",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "marsanben92@gmail.com",
            "name": "Marcos",
            "username": "sonirico"
          },
          "committer": {
            "email": "marsanben92@gmail.com",
            "name": "Marcos",
            "username": "sonirico"
          },
          "distinct": true,
          "id": "f91953a63a111eaeeb38e64b856ec959cf9c9b36",
          "message": "chore: relax linter rules to focus on critical issues\n\nDisabled noisy linters that don't catch bugs:\n- unused: Too noisy for private fields reserved for future use\n- unconvert: Style issue, not a bug\n- unparam: Too noisy\n- prealloc: Micro-optimization\n- fieldalignment: Micro-optimization, not worth the noise\n\nDisabled gocritic style checks:\n- unnamedResult, sloppyReassign, unlambda\n\nAdded gosimple to test file exclusions.\n\nThis focuses the linter on actual bugs (errcheck, copylocks, staticcheck)\nwhile reducing noise from style preferences.",
          "timestamp": "2026-02-07T18:37:45+01:00",
          "tree_id": "86085071be98f0e35de9ceb9a6950c73a2e76c43",
          "url": "https://github.com/sonirico/parco/commit/f91953a63a111eaeeb38e64b856ec959cf9c9b36"
        },
        "date": 1770486302219,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1920,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "591492 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1920,
            "unit": "ns/op",
            "extra": "591492 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "591492 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "591492 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "591492 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 13325,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "91632 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 13325,
            "unit": "ns/op",
            "extra": "91632 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "91632 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "91632 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "91632 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 132515,
            "unit": "ns/op\t      8123 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "9320 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 132515,
            "unit": "ns/op",
            "extra": "9320 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "9320 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "9320 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9320 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1746,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "637875 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1746,
            "unit": "ns/op",
            "extra": "637875 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "637875 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "637875 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "637875 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 12243,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "91646 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 12243,
            "unit": "ns/op",
            "extra": "91646 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "91646 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "91646 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "91646 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 115890,
            "unit": "ns/op\t      8123 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "9752 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 115890,
            "unit": "ns/op",
            "extra": "9752 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "9752 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "9752 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9752 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2531,
            "unit": "ns/op\t       270.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "449986 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2531,
            "unit": "ns/op",
            "extra": "449986 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 270,
            "unit": "payload_bytes/op",
            "extra": "449986 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "449986 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "449986 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 26367,
            "unit": "ns/op\t      1678 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "45112 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 26367,
            "unit": "ns/op",
            "extra": "45112 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1678,
            "unit": "payload_bytes/op",
            "extra": "45112 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "45112 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "45112 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 318826,
            "unit": "ns/op\t     16609 payload_bytes/op\t   76511 B/op\t    2003 allocs/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 318826,
            "unit": "ns/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16609,
            "unit": "payload_bytes/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76511,
            "unit": "B/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3583 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2542,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "429465 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2542,
            "unit": "ns/op",
            "extra": "429465 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "429465 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "429465 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "429465 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 19336,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "61784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 19336,
            "unit": "ns/op",
            "extra": "61784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "61784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "61784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "61784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 188768,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "5941 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 188768,
            "unit": "ns/op",
            "extra": "5941 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "5941 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "5941 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "5941 times\n4 procs"
          }
        ]
      }
    ]
  }
}