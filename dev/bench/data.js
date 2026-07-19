window.BENCHMARK_DATA = {
  "lastUpdate": 1784467665023,
  "repoUrl": "https://github.com/sonirico/parco",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "marsanben92@gmail.com",
            "name": "Marcos Benedicto",
            "username": "sonirico"
          },
          "committer": {
            "email": "marsanben92@gmail.com",
            "name": "Marcos Benedicto",
            "username": "sonirico"
          },
          "distinct": true,
          "id": "0a268309f252c970278357e495ca100f31435caa",
          "message": "docs: require Go 1.22+",
          "timestamp": "2026-07-19T14:56:52+02:00",
          "tree_id": "b4a848e6c26ff0d5117eb44f3cca5a2878ab7172",
          "url": "https://github.com/sonirico/parco/commit/0a268309f252c970278357e495ca100f31435caa"
        },
        "date": 1784467664763,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1204,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "952974 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1204,
            "unit": "ns/op",
            "extra": "952974 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "952974 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "952974 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "952974 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 8138,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "145785 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 8138,
            "unit": "ns/op",
            "extra": "145785 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "145785 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "145785 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "145785 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 77213,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15580 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 77213,
            "unit": "ns/op",
            "extra": "15580 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15580 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15580 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15580 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1242,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "946392 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1242,
            "unit": "ns/op",
            "extra": "946392 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "946392 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "946392 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "946392 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 8141,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "146169 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 8141,
            "unit": "ns/op",
            "extra": "146169 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "146169 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "146169 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "146169 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 76740,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 76740,
            "unit": "ns/op",
            "extra": "15432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15432 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2456,
            "unit": "ns/op\t       271.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "460720 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2456,
            "unit": "ns/op",
            "extra": "460720 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 271,
            "unit": "payload_bytes/op",
            "extra": "460720 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "460720 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "460720 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 24686,
            "unit": "ns/op\t      1678 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "48262 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 24686,
            "unit": "ns/op",
            "extra": "48262 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1678,
            "unit": "payload_bytes/op",
            "extra": "48262 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "48262 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "48262 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 315493,
            "unit": "ns/op\t     16583 payload_bytes/op\t   76499 B/op\t    2003 allocs/op",
            "extra": "3693 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 315493,
            "unit": "ns/op",
            "extra": "3693 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16583,
            "unit": "payload_bytes/op",
            "extra": "3693 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76499,
            "unit": "B/op",
            "extra": "3693 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3693 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2274,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "487221 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2274,
            "unit": "ns/op",
            "extra": "487221 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "487221 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "487221 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "487221 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 16258,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "72784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 16258,
            "unit": "ns/op",
            "extra": "72784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "72784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "72784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "72784 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 160056,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "6292 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 160056,
            "unit": "ns/op",
            "extra": "6292 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "6292 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "6292 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "6292 times\n4 procs"
          }
        ]
      }
    ]
  }
}