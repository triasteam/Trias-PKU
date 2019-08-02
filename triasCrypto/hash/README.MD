hash方法 | 功能测试结果  
-|-
Md5         |    fc059406246bdfebbfe83d825ead604e   
sha1        |    674ecfadc54b0fa78663f3f57590bbd351b0a5ea   
sha256224   |    7eeb165aa433f8afd971d01962bff48dcfc2c290788952004ed7f170   
sha256256   |    43e1a9af70daf9d1eaf3782e26f8f8324ff94c509ecb34287a946170ac96be6a   
sha512224   |    b372375c2c7289507dd50eaf401a3599bd9b92c36b264ef138747c31   
sha512256   |    ce5e291c8a1d254b515c7a2c5c2a4fe00ef1f6e487c38e88b2b8cfc11c5bf39d   
sha512384   |    15e5790b4b949923de2f1c509c478cc85d195bde0b326b580eacff15efb98f0458bfd09643aceeaa6d5166cef708850b   
sha512512   |    f75bf1dece1a88b167fbc9b27b5f81796b632f4115c325d808316f3f82db29d765a9154704879fdf22f66cd741e32d70ccecf293a87ff7d6d2a3486791d0a99f   
TriasRipemd160 | f06972c839fdb1d50500810403f707c4589271d0   
TriasSHA3_224  | d7ad4687bdcf458a19a79cb9c6c409af5a85e0387b86aecfc5aedaee   
TriasSHA3_256  | 83466c996e7b9b5320e2045fde6498dab502b96bad0b67e321596f766cc61b05   
TriasSHA3_384  | 6e8f41a8bd323ddb34074cc4360c85b854b78d6fde61d81a0227161782aa00036bfb8ed24cb89b0861e44d062ac68f05   
TriasSHA3_512  | 2a2a265877021e31796aa86891c79ade5b677b0fa5b95fb80f1d39a31274aea11a15f2b3fa02a9801f2e9d62a3e35532b3edb74f24fce7c6204b7b532e467a07   
TriasBlake2s_256 | 0e3647b05932f6907cf6e0ffd2d56514f3d6f093f217f4c092f784c8907ed9ba   
TriasBlake2b_256 | a2d558398e6c3465312a38fe56bd570f0659b417c2da8d38bdd5d0f2652560f8   
TriasBlake2b_384 | ba61e8bd6d1b494f494120be9f7ab7903c25930d06715e01c16c350bbfe74d4f49e83efc78d4a5f4c528b525fea92a6f   
TriasBlake2b_512 | c99b03d4b829ffbaac218a26f16575661d2dd09e8d45efcaff09346557df19303eb75140c86724010cd52878bc95d4c68f11941dbadba39ecf57019544fef9ce   

hash方法 | 测试次数 | 性能测试结果   
-|-|-
BenchmarkTriasMd5           |     5000000      |         284 ns/op   
BenchmarkTriasSha1          |     5000000      |         335 ns/op    
BenchmarkTriasSha224        |     3000000      |         422 ns/op   
BenchmarkTriasSha256        |     3000000      |         431 ns/op   
BenchmarkTriasSha512224     |     3000000      |         515 ns/op   
BenchmarkTriasSha512256     |     3000000      |         514 ns/op   
BenchmarkTriasSha512384     |     3000000      |         544 ns/op   
BenchmarkTriasSha512512     |     3000000      |         584 ns/op   
BenchmarkTriasRipemd160     |     2000000      |         734 ns/op   
BenchmarkTriasSHA3_224      |     2000000      |         973 ns/op   
BenchmarkTriasSHA3_256      |     2000000      |         978 ns/op   
BenchmarkTriasSHA3_384      |     2000000      |         957 ns/op   
BenchmarkTriasSHA3_512      |     2000000      |         955 ns/op   
BenchmarkTriasBlake2s_256   |     5000000      |         349 ns/op   
BenchmarkTriasBlake2b_256   |     3000000      |         404 ns/op   
BenchmarkTriasBlake2b_384   |     3000000      |         436 ns/op   
BenchmarkTriasBlake2b_512   |     3000000      |         462 ns/op   
