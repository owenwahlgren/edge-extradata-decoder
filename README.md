
# Polygon Edge 0.6.0 (BLS Update) extraData Decoder

Decode the extraData parameter in block headers from Polygon Edge
0.6.0 (BLS)

## Build the project


`$ go build decoder.go`

## Example

Input

`$ ./decoder 0x0000000000000000000000000000000000000000000000000000000000000000f901e7f8d8f8469456fae64bde3d31a884f699fe88ff5104c71d61a6b0b34bc758246994f6d86a7715d81e359fd516634430c16e848480667afeb0de51f2c5b9b69b1550c2f1b83af36f8b92e0f84694b53fd6be68693e33dceda32722de7c1d0e26f618b08976336e55678b8bfb8c5deabc240be96a32be9f9b443803825a37444bc9234c62029e376add89c8349fcf7e7674d7ecf846944f944c7fc8377c6b5dd871d95fa4b7f8ff889fa6b0a7bad64c86311303b671537155c3d7a81dd5a29ebf3298c6a5c72ababe06a2bb51c9f1ba411016359c17cb789d533045b841b87d9f72c07297850bdfa046e924829bc4af2f36a6527640abb6de382876eb8261255ba21b83a09b17eb2a5c076fca2e0a503a760c6b5635eacec1da7d79986300f86307b860a22b2b01ba85cef2c5c8ec65c975f0edd24b4670c50789a7f663ad3c103d456e02ff386a46ce8725453d185744b3efa30f1ac979e2530ea01beb26a1c5308065975ac2317264b5b28a106805226d985f84586f4779b5838cfa014d589edb20f4f86307b860a03f3c439a02354317d2d9575ea2f93d7aba84e64db003c50dadef0ad5e4655b180f1ed33402bcd16193e3ebe3ec9a761773a61ba4027ed9b8cd93109fb9590f44b2881ebac0deec6e16e9c397f57f12f66b2e585dd4d515b55653d6b10754c5`

Output

`validators: 3:	address:bls`

`validator: 0x56FAe64bDE3d31a884f699Fe88Ff5104C71d61a6:0xb34bc758246994f6d86a7715d81e359fd516634430c16e848480667afeb0de51f2c5b9b69b1550c2f1b83af36f8b92e0`

`validator: 0xb53Fd6bE68693E33DCeDA32722de7c1d0e26F618:0x8976336e55678b8bfb8c5deabc240be96a32be9f9b443803825a37444bc9234c62029e376add89c8349fcf7e7674d7ec`

`validator: 0x4F944c7fc8377c6B5DD871D95fa4B7f8ff889fa6:0xa7bad64c86311303b671537155c3d7a81dd5a29ebf3298c6a5c72ababe06a2bb51c9f1ba411016359c17cb789d533045`

