window.BENCHMARK_DATA = {
  "lastUpdate": 1792340705040,
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
      },
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
          "id": "82bcd1e3163e451e0443ffa60be935cbb90c5804",
          "message": "fix: unbreak CI lint and Go 1.22 macOS job\n\n- replace deprecated reflect.StringHeader/SliceHeader (SA1019) with\n  unsafe.String/unsafe.Slice in String2Bytes/Bytes2String\n- exclude Go 1.22 on macos-latest: its binaries lack the LC_UUID load\n  command required by the dyld on current runners\n- pin golangci-lint to v1.64.8 in both CI and 'just setup' so local\n  lint reproduces CI exactly\n- swap shields.io badges (license, go version) for badgen.net: the\n  GitHub-backed shields endpoints are returning 520s",
          "timestamp": "2026-07-19T15:38:01+02:00",
          "tree_id": "0b1437a8872fad392e4643884326c812b8c9238b",
          "url": "https://github.com/sonirico/parco/commit/82bcd1e3163e451e0443ffa60be935cbb90c5804"
        },
        "date": 1784468577612,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1180,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "984440 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1180,
            "unit": "ns/op",
            "extra": "984440 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "984440 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "984440 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "984440 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 8088,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "135432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 8088,
            "unit": "ns/op",
            "extra": "135432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "135432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "135432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "135432 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 76053,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15898 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 76053,
            "unit": "ns/op",
            "extra": "15898 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15898 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15898 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15898 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1182,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "950554 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1182,
            "unit": "ns/op",
            "extra": "950554 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "950554 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "950554 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "950554 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 8024,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "148645 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 8024,
            "unit": "ns/op",
            "extra": "148645 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "148645 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "148645 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "148645 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 75525,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15957 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 75525,
            "unit": "ns/op",
            "extra": "15957 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15957 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15957 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15957 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2243,
            "unit": "ns/op\t       268.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "501798 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2243,
            "unit": "ns/op",
            "extra": "501798 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 268,
            "unit": "payload_bytes/op",
            "extra": "501798 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "501798 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "501798 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 22859,
            "unit": "ns/op\t      1668 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "52365 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 22859,
            "unit": "ns/op",
            "extra": "52365 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1668,
            "unit": "payload_bytes/op",
            "extra": "52365 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "52365 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "52365 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 307194,
            "unit": "ns/op\t     16647 payload_bytes/op\t   76499 B/op\t    2003 allocs/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 307194,
            "unit": "ns/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16647,
            "unit": "payload_bytes/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76499,
            "unit": "B/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2126,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "523765 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2126,
            "unit": "ns/op",
            "extra": "523765 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "523765 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "523765 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "523765 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 16105,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "75132 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 16105,
            "unit": "ns/op",
            "extra": "75132 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "75132 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "75132 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "75132 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 153182,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 153182,
            "unit": "ns/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "6961 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sonirico@users.noreply.github.com",
            "name": "Marquitos",
            "username": "sonirico"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ad4e9301a9113026e75ba2a20d4e0f4da609dc7f",
          "message": "Merge pull request #5 from sonirico/experiment/generic-methods\n\nfeat: Model builder, Registry and parco-gen behind go1.27 (v0.14.0)",
          "timestamp": "2026-10-18T14:07:51+02:00",
          "tree_id": "b8cf1865ab70532d7a9fda59919edabb454ecb95",
          "url": "https://github.com/sonirico/parco/commit/ad4e9301a9113026e75ba2a20d4e0f4da609dc7f"
        },
        "date": 1792325321424,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1246,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1246,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 7619,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "153508 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 7619,
            "unit": "ns/op",
            "extra": "153508 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "153508 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "153508 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "153508 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 73490,
            "unit": "ns/op\t      8123 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "16312 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 73490,
            "unit": "ns/op",
            "extra": "16312 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "16312 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "16312 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16312 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1138,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1138,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 7557,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "157471 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 7557,
            "unit": "ns/op",
            "extra": "157471 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "157471 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "157471 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "157471 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 72711,
            "unit": "ns/op\t      8123 payload_bytes/op\t     185 B/op\t       3 allocs/op",
            "extra": "16771 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 72711,
            "unit": "ns/op",
            "extra": "16771 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "16771 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "16771 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16771 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2401,
            "unit": "ns/op\t       269.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "467160 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2401,
            "unit": "ns/op",
            "extra": "467160 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 269,
            "unit": "payload_bytes/op",
            "extra": "467160 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "467160 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "467160 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 25824,
            "unit": "ns/op\t      1673 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "47503 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 25824,
            "unit": "ns/op",
            "extra": "47503 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1673,
            "unit": "payload_bytes/op",
            "extra": "47503 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "47503 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "47503 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 314182,
            "unit": "ns/op\t     16595 payload_bytes/op\t   76499 B/op\t    2003 allocs/op",
            "extra": "3562 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 314182,
            "unit": "ns/op",
            "extra": "3562 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16595,
            "unit": "payload_bytes/op",
            "extra": "3562 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76499,
            "unit": "B/op",
            "extra": "3562 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3562 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2311,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "498123 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2311,
            "unit": "ns/op",
            "extra": "498123 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "498123 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "498123 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "498123 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 16655,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "71553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 16655,
            "unit": "ns/op",
            "extra": "71553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "71553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "71553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "71553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 161713,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "6560 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 161713,
            "unit": "ns/op",
            "extra": "6560 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "6560 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "6560 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "6560 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sonirico@users.noreply.github.com",
            "name": "Marquitos",
            "username": "sonirico"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1a596bdc7ed0c377aa12a03903eba4d259d3dfab",
          "message": "Merge pull request #6 from sonirico/fix/ci-lint-and-badges\n\nfix: green CI after v0.14.0 and replace dead README badges",
          "timestamp": "2026-10-18T17:19:36+02:00",
          "tree_id": "f8b861ce3e4293bcba7263612f0c26b83c9b8a9c",
          "url": "https://github.com/sonirico/parco/commit/1a596bdc7ed0c377aa12a03903eba4d259d3dfab"
        },
        "date": 1792336819285,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1140,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "993946 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1140,
            "unit": "ns/op",
            "extra": "993946 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "993946 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "993946 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "993946 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 7474,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "160520 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 7474,
            "unit": "ns/op",
            "extra": "160520 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "160520 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "160520 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "160520 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 71286,
            "unit": "ns/op\t      8123 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "16744 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 71286,
            "unit": "ns/op",
            "extra": "16744 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "16744 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "16744 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16744 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1132,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1132,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 7476,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "159042 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 7476,
            "unit": "ns/op",
            "extra": "159042 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "159042 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "159042 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "159042 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 71232,
            "unit": "ns/op\t      8123 payload_bytes/op\t     185 B/op\t       3 allocs/op",
            "extra": "16593 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 71232,
            "unit": "ns/op",
            "extra": "16593 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "16593 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "16593 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16593 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2405,
            "unit": "ns/op\t       268.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "476944 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2405,
            "unit": "ns/op",
            "extra": "476944 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 268,
            "unit": "payload_bytes/op",
            "extra": "476944 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "476944 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "476944 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 24713,
            "unit": "ns/op\t      1678 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "48510 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 24713,
            "unit": "ns/op",
            "extra": "48510 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1678,
            "unit": "payload_bytes/op",
            "extra": "48510 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "48510 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "48510 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 312021,
            "unit": "ns/op\t     16605 payload_bytes/op\t   76511 B/op\t    2003 allocs/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 312021,
            "unit": "ns/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16605,
            "unit": "payload_bytes/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76511,
            "unit": "B/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3429 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2283,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "475082 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2283,
            "unit": "ns/op",
            "extra": "475082 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "475082 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "475082 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "475082 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 16552,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "71212 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 16552,
            "unit": "ns/op",
            "extra": "71212 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "71212 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "71212 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "71212 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 162140,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 162140,
            "unit": "ns/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "6642 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "6642 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "sonirico@users.noreply.github.com",
            "name": "Marquitos",
            "username": "sonirico"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c1aa1f0ae88d26b170469ec926c6bcc6586d9c6c",
          "message": "Merge pull request #7 from sonirico/perf/parse-reader\n\nperf: one pooled parseReader per Parse instead of a pooled buffer per field",
          "timestamp": "2026-10-18T18:24:13+02:00",
          "tree_id": "e7ab21c90ff57014020cb7885214371524db4f60",
          "url": "https://github.com/sonirico/parco/commit/c1aa1f0ae88d26b170469ec926c6bcc6586d9c6c"
        },
        "date": 1792340704705,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size",
            "value": 1226,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "916285 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - ns/op",
            "value": 1226,
            "unit": "ns/op",
            "extra": "916285 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "916285 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "916285 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "916285 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size",
            "value": 8398,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "143031 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - ns/op",
            "value": 8398,
            "unit": "ns/op",
            "extra": "143031 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "143031 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "143031 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "143031 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size",
            "value": 78576,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15274 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - ns/op",
            "value": 78576,
            "unit": "ns/op",
            "extra": "15274 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15274 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15274 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoAlloc_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15274 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size",
            "value": 1215,
            "unit": "ns/op\t        91.00 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "947614 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - ns/op",
            "value": 1215,
            "unit": "ns/op",
            "extra": "947614 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - payload_bytes/op",
            "value": 91,
            "unit": "payload_bytes/op",
            "extra": "947614 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "947614 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/small_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "947614 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size",
            "value": 8314,
            "unit": "ns/op\t       742.0 payload_bytes/op\t     184 B/op\t       3 allocs/op",
            "extra": "142815 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - ns/op",
            "value": 8314,
            "unit": "ns/op",
            "extra": "142815 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - payload_bytes/op",
            "value": 742,
            "unit": "payload_bytes/op",
            "extra": "142815 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "142815 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/medium_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "142815 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size",
            "value": 81636,
            "unit": "ns/op\t      8123 payload_bytes/op\t     186 B/op\t       3 allocs/op",
            "extra": "15332 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - ns/op",
            "value": 81636,
            "unit": "ns/op",
            "extra": "15332 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - payload_bytes/op",
            "value": 8123,
            "unit": "payload_bytes/op",
            "extra": "15332 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - B/op",
            "value": 186,
            "unit": "B/op",
            "extra": "15332 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoDiscard_Compile/large_size - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15332 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size",
            "value": 2509,
            "unit": "ns/op\t       267.0 payload_bytes/op\t     970 B/op\t      23 allocs/op",
            "extra": "476839 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - ns/op",
            "value": 2509,
            "unit": "ns/op",
            "extra": "476839 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - payload_bytes/op",
            "value": 267,
            "unit": "payload_bytes/op",
            "extra": "476839 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - B/op",
            "value": 970,
            "unit": "B/op",
            "extra": "476839 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/small_size - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "476839 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size",
            "value": 24922,
            "unit": "ns/op\t      1677 payload_bytes/op\t    7685 B/op\t     203 allocs/op",
            "extra": "47964 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - ns/op",
            "value": 24922,
            "unit": "ns/op",
            "extra": "47964 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - payload_bytes/op",
            "value": 1677,
            "unit": "payload_bytes/op",
            "extra": "47964 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "47964 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/medium_size - allocs/op",
            "value": 203,
            "unit": "allocs/op",
            "extra": "47964 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size",
            "value": 313668,
            "unit": "ns/op\t     16629 payload_bytes/op\t   76499 B/op\t    2003 allocs/op",
            "extra": "3553 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - ns/op",
            "value": 313668,
            "unit": "ns/op",
            "extra": "3553 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - payload_bytes/op",
            "value": 16629,
            "unit": "payload_bytes/op",
            "extra": "3553 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - B/op",
            "value": 76499,
            "unit": "B/op",
            "extra": "3553 times\n4 procs"
          },
          {
            "name": "BenchmarkJson_Compile/large_size - allocs/op",
            "value": 2003,
            "unit": "allocs/op",
            "extra": "3553 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size",
            "value": 2303,
            "unit": "ns/op\t       155.0 payload_bytes/op\t     762 B/op\t      25 allocs/op",
            "extra": "494830 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - ns/op",
            "value": 2303,
            "unit": "ns/op",
            "extra": "494830 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - payload_bytes/op",
            "value": 155,
            "unit": "payload_bytes/op",
            "extra": "494830 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - B/op",
            "value": 762,
            "unit": "B/op",
            "extra": "494830 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/small_size - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "494830 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size",
            "value": 16569,
            "unit": "ns/op\t       991.0 payload_bytes/op\t    4068 B/op\t     207 allocs/op",
            "extra": "71974 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - ns/op",
            "value": 16569,
            "unit": "ns/op",
            "extra": "71974 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - payload_bytes/op",
            "value": 991,
            "unit": "payload_bytes/op",
            "extra": "71974 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - B/op",
            "value": 4068,
            "unit": "B/op",
            "extra": "71974 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/medium_size - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "71974 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size",
            "value": 159286,
            "unit": "ns/op\t     10171 payload_bytes/op\t   37437 B/op\t    2007 allocs/op",
            "extra": "6885 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - ns/op",
            "value": 159286,
            "unit": "ns/op",
            "extra": "6885 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - payload_bytes/op",
            "value": 10171,
            "unit": "payload_bytes/op",
            "extra": "6885 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - B/op",
            "value": 37437,
            "unit": "B/op",
            "extra": "6885 times\n4 procs"
          },
          {
            "name": "BenchmarkMsgpack_Compile/large_size - allocs/op",
            "value": 2007,
            "unit": "allocs/op",
            "extra": "6885 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/small_size",
            "value": 1383,
            "unit": "ns/op\t     680 B/op\t       9 allocs/op",
            "extra": "836437 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/small_size - ns/op",
            "value": 1383,
            "unit": "ns/op",
            "extra": "836437 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/small_size - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "836437 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/small_size - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "836437 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/medium_size",
            "value": 9063,
            "unit": "ns/op\t    4192 B/op\t      99 allocs/op",
            "extra": "130796 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/medium_size - ns/op",
            "value": 9063,
            "unit": "ns/op",
            "extra": "130796 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/medium_size - B/op",
            "value": 4192,
            "unit": "B/op",
            "extra": "130796 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/medium_size - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "130796 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/large_size",
            "value": 93252,
            "unit": "ns/op\t   61249 B/op\t    1001 allocs/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/large_size - ns/op",
            "value": 93252,
            "unit": "ns/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/large_size - B/op",
            "value": 61249,
            "unit": "B/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoBytes_Parse/large_size - allocs/op",
            "value": 1001,
            "unit": "allocs/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/small_size",
            "value": 2124,
            "unit": "ns/op\t     680 B/op\t       9 allocs/op",
            "extra": "543018 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/small_size - ns/op",
            "value": 2124,
            "unit": "ns/op",
            "extra": "543018 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/small_size - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "543018 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/small_size - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "543018 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/medium_size",
            "value": 15093,
            "unit": "ns/op\t    4194 B/op\t      99 allocs/op",
            "extra": "78810 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/medium_size - ns/op",
            "value": 15093,
            "unit": "ns/op",
            "extra": "78810 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/medium_size - B/op",
            "value": 4194,
            "unit": "B/op",
            "extra": "78810 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/medium_size - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "78810 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/large_size",
            "value": 152756,
            "unit": "ns/op\t   61350 B/op\t    1001 allocs/op",
            "extra": "7797 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/large_size - ns/op",
            "value": 152756,
            "unit": "ns/op",
            "extra": "7797 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/large_size - B/op",
            "value": 61350,
            "unit": "B/op",
            "extra": "7797 times\n4 procs"
          },
          {
            "name": "BenchmarkParcoReader_Parse/large_size - allocs/op",
            "value": 1001,
            "unit": "allocs/op",
            "extra": "7797 times\n4 procs"
          }
        ]
      }
    ]
  }
}