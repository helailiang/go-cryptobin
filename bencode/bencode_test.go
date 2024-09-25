package bencode

import (
    "fmt"
    "time"
    "bytes"
    "testing"
    "encoding/hex"

    cryptobin_test "github.com/deatil/go-cryptobin/tool/test"
)

var torrentData = "64383a616e6e6f756e636534323a687474703a2f2f747261636b65722e617263686c696e75782e6f72673a363936392f616e6e6f756e6365373a636f6d6d656e7434313a41726368204c696e757820323031312e30382e313920287777772e617263686c696e75782e6f72672931303a6372656174656420627931333a6d6b746f7272656e7420312e3031333a6372656174696f6e2064617465693133313337353031373165343a696e666f64363a6c656e6774686931383937393232353665343a6e616d6534303a617263686c696e75782d323031312e30382e31392d6e6574696e7374616c6c2d693638362e69736f31323a7069656365206c656e6774686935323432383865363a706965636573373234303aff5d9661ce75bc30c3311311f6e4d1546c6b20a5f0275594c137a0d3270fb8d9f6553b9b7cb771af8d2c4c50fcb6bb9f04ee27c917d56787588927f7774708d811a133c86fd38384917e04eb3335ce582cf6d41fc58369ca6b831aae9b4b3188fe0d6b2a941bc293ebd94112dd9e9d0b69456aec6d5af6bafeacb61ed5c0917c019bc5a46afda1f1ec47581e66d80c5c001c7702b42f6cb4d0b9c042c8133a51e7f1816de66ad5007e26f113828b8d8393c8a13716085009659053d43477686c3046e0de53dd3735f093e79ef8531930cbe4d14354681bd14d7d03553331c15dc340c4bea864cdfabc214be193ffc9a9b466862d0a41b1a363fcb0b1d14fa3b15c5eb80dbe1d42c42dae1fc889d3108ba7483e1045ccaeacfea8e20113bd74596d4a1f19237d1f4a0c9ad95af08ee5d5abbabdf9e4b17389cce99373ffb3e6d64e5d73bd0be56aad796ba842c31b35e64fbb3f3f8b27f009aa772c984b91bf0691f37a64d92b48f1408230c1626502e2e88d0c758d25faae1fed97048af69e602f4867aa5f08f44cbc6ed40179b729f0e0df2aadea172332840266572fe4ff4ce1a71148624462650a7cd2fbc73dc652a984c1baaadc89dc11f24de7c88bd00e71491563d7d03db0adadf37a50e039af3f7843ebf4db2f81fbd364c2b44c574a5fce38a3d8b5a808208e0ea3e4c659721ef7f5d2f6658c06a6d01908e0ee0298c030efe9a7e9008bba413a5682e2df6744ff2a45b756fee9a7c8f096aa8637c35fcc161edcbdafab9af9105a18edf37618cf6779c93099aeae92e5143e5ffc9ef392f9964ce7a6b8910dcabcf65b1b1de03398da6cf3368cf2eee01d36844619d778dc8c0ff5f0df06c6788cb966c978533783825b7c243d63f81126257dc43d82b9a02d7471a4a92ac5f93fd87193a9c73ae4c25d8d16662ed4f30b32d922e625e05db2100a92c1710a29cacdde152b0606de308fc8ab8ebf3465a7408c452fc016541f5f9b400523d570b748796a018ada0c73b4266cad09d94bf9bc15eea47ea428cb4ea8f9d62ad146a0a6f5a2a7c497743445f6715e60a36a2d6659b42065c8f1d7a22218ace81b194afdcc38d8e62263ecd8e2b59fd2467d2a0dde5d2bfe2fbdb64061ce210179deef40337f7d93afe3974ee60cb1e1e40c96007df388babb293bbb30fcffc70e5c9b3a76cf881155304b2085b822c3cc6340cc0c1902fb4b40a763c1fab6b4603a05e95d91a55bd757ff80dcea2d683884985131723f681e23f7d9a91cef65a523c0a3b7c6dcb7b60896bfc367748baba2c107a7667ebfad38b711663f6353484bbc75c935f662895199620be3d1f95fdaf11012bc69fda49b794ce3a19f276496da4938695e5e9e00e9a9b40fcacb21de58b55666e2cb8ae49346b61e84740905dfcd6096764c6986fe594d2fc41c6c36d00e1c2c5b99712ffd33ad20687f6cd0152a086b30f1571b8d58331d49fd8725b38131c10ea944cd0604cb896972c35b333d7fdfffc24c2db376f206d73ada5128d6a567c45b7b2bdb85a252f8751547d1ae76efc1962212dc8b2c58a30181bc5b309654d9f1f4bf2df583810787be56f23fccb296f438ac3c52706b945fa4ac69d3bed7458f2986cbf8fb5610c84f430b837f47aed2fc3ba22bfd6ceeb4d2752ba9df4b4c61566ef01849aaa9aa8483af385ac06bdfdb490d941de17a470f30382c17cf430c2d4f6b6368d68449dfa892c8428bf9c1601c57036b5422a21a324943e1e0616cb1c294972af4d351623a22a9742b1d218a6f968e2e149c60640dd9e14dcb4bee24a1955c7afb3bbaec2ca9bb37a69a927d72d499743a2c7824ce627cfe764663e1ef7dcf3f1790684c42df35c03ef4732d2d899eec74a448d4b284c43fd79dbef830005511a0ca164b2366373a78a3fa8871460b58bb1211257e84d4c409517a65b5efc3c898a687f49fbec3fdca7c675ca39895d05362c8ef0ff151c823c1974606c2aa3550e0356d426601cb93103e6c6ea0d665f2345ca63dd133b74cb92e05cc066f497a3d026ee728c7b74f53fa104a2685546e71efc0bbb0ce56fc92620a3b6380a62cc295e5bd68470d729ef5ca90967623c5fdc4216c813b9e6178b4e851729502d5f7161fa3cbf7a2184f37de42cee31fe49cc4f900c92eb40e02d5650ed43940d01760db49efc58b9525f64641427dffac5821647f719daebac86f07192e9e5fee096e28fe69ba3c188b44c56cfebd02797282ad7557f4835784beec06958b6ab9fbf0c1c99d181036affcb110be7dd711068289a2f4863859af97d310fa699937383f282f0df3c4c560d0d4e69ccd8e419a2666fa7a66696e07657f0b5536c2634ff6ed549d22d46f8b852a6ed79baaacbbfd92b8aacc1e88e5351f148516e2a25aad6cb4779bcc5ad16d5021335327aae1abf7d89f25b3b7359d153b2a2cac76d03bdf156a32ecdc1c3c5941dfc92acb5c6f108b15a4e632a828fa2658bc86b06d05123550760a3fa46a6ba30693f7a71fcf0b29ddd24af4ccb5f531e53b8b919f71e3249264eb15a3943245a805a51e027c41f1788163e2341609facc01a4915b34c0f13e30749d1779a4e005c618aee4274c3e93116f7f1e397493ba0d2cd4f06e790a9273e8ec3760bf56d58439993205673dfd55c5e8790539e1dbe17c0c297b8711c6813e9eff9c9f065aa9103469e9a5de4f7d12af9a0cc986b99242e0160d14aa8a605290c340e1335e8ac22214b80d35ff5053f87c9718e1b8acfe59fc92c9d03fc0e2aac81e8c1b93cdaaf0ce6dd57c8bd18714f021102e4db7c17dea334e9568b4a11be31889ed0360cb3eb1653d680e06f2832822231aebb103d8408bf01c436682c8d47fe42f5f0a5296239b53f14c9349dabde7f3fd394dd05f4d49e9d61c2ab23b00212fbc5578b99b0a8ac3f4610b4e434a15a091c50dc8f592d7cc33313476dc850731d00b9c569a8e40086cb65d0b653d303e94725b7b23cb429a9b608bfb4090ed16efb3440e3d933a06e95226b3bf929e91d3b07120752ccecd5bf68dfa559b7dcafe6acf9fb4ace4b524107a46f1e8678a11753f552c508bd618c17865833a6cbf0f11f4af8f0647c3772c01da5616d0c91c697acd20deeb4aa5324613f2ac90aec9a15406f43059ac0c1d5cee32ae3606ef8edd75cd9891f5cd405181935288be8c549cf09c0bbbf32a95d9e2809ceee2b5beb54bee92df87dcafc65b17f61fa3f969feb616bbb9650a71219c214ad938b5877e05de75c116786eaf788372e341860e9cd26f0e06d164b57d4ecf349c39e7a5d23ce2bc53a98cb5d830e26fecdd570418a3d7db8c3540e76812784b812924f8f065f860e1aa46db56034b6520f9bf8f2a238d003d4758eb5116790942d33454548816457a61ff3b7ba1fc67725ef95684ac5e150fcec54dee910077642b9fa85d999553a147c53b9d566705a1733399d48be4ceff9fad1e7734ac6c973c39940faf93d35cf8aafffa6710d676d5a16f60871f59584c4fd9a87ce0c1d2f74dcf2b36c79bdf4c1ed468d8f8df3a1014e02f64d73e12a96e86c2586e6d0a0e9e6f230b4bd1c3ad0ba44326c31f4872969877b053770b24d30cd3b9c77ccfed7f9ab2e10419ebd8027323671c463f981abe9e90fd81b50f68abee758e4635edf16bbd5528f0ee5169cbd5d61503079950aefde0c7ae7643ebee523d287861dd2e567d6e5a0ef4c2a0a752f369906976872d5b993cc0040cba41592ec4544af5bbf1f0d7c656fd3877a3b080bc9b6a419dcf62e4fea606a4e0f8de01af38cc5dd60a2014c5fb7f89c093aa8163722eafa96134bc14de0dbcc52b7e9778d604936742bfd1ed4ac46f8873c3d1d0c5acc0a0f26e91279619fd77c7fce294e97e0e8ad7a97f4313d5a60b3f901652d5f8386084db01c1db38dfc065e8b2ffb04b1ce91d1b7966fdbd283d389c2d898a4cf4c481ccbc1256bdb82e905b71eeadb0319a3ddaabb169c50269def0731c3d7fb263efbd12a6b752aad69ab241da7ad576c581c456d85c0ad732928b3f70f84c08f01654dd9d56d410c997fe42d9e18802b1b549233d3e2dab9774b26a18bbc7fc1d12b4846ef69ea5bea81a6aab0d646de3155008f43d9b3674db0e31862b4f9ff9324d10fd41031e4e1070188e9efff1704f4b93c8a0f6b0ac0a71e5d1929155eecefe14f36fa7ed8dd8296c4522972f63f89c19f5186672ca8fdee0b448a5fd04a887734ff1b5bc3e2df5d4527e4c6a330179cb60af17237eadf65b38906ae61c7b59bbb28c9674c35af0d6e8f32d0036a843e999b28c4ee0f8ce2e278fbaef96808e402914768b95565b2c1d486f35868b7125264e51e29b0e49330650699789c0b897d2d736260d2a493fc168f9f83fb04cbd6081d9ab52ad57b2108769626a3c5d8a42f47250c18e6e19affe3fe5a6ae6f411bb3555b25f9d50c0a8e6cf491deb962b42f1819f4d63fe700887874e43de20dc783c805c978f1db6b7b63c5c02222d27793a128aac5e3bc1fe780ad17efeffbed79245540a1354d42b925f1e1e39f5899d2720485d564b14afd4976fe2148e05c7e556af1c8eecad4543372d2e6a3b5e7ffd1d778d8ffc6a9c9fd8b028e267ffa1b4047e78fb900038f2657efc45b1d637f4049a9b6cb1a7d96c3d87cac2a6af91951b8d2ab0441f7fcf88ea1a831f0fb56fc387637ba21aa80ad2a0d6e7645160ff20d28a0125cafbfc207e6a75ab170a30ac0f8dee941b095aa65b8461a38b77a97c384eaf6274d7c09d80f042eea521e973dae7376e53283b622e84412cd2e3a21a8732ce9b7ca07bcbb555e58e973666af165fa097df0bc30a736a5f2c5870f950bd3d27e5b4f0d48ae7d41b1f807758a19b3486703f319f845e5aa2914d2feb328c6b26c2df8dfccd636ba556ab42b13c559f5c66141953778f640b7e5f6e08ac90b8fc8a9584a95720018446003ddf6c02ec18fe3f4e86cc24e55096515bdb44b4bca624dd37212dafb21dcca57a729299b2e101311a31d900f33bdcbab3db001b0735b56963199f7ffb40ea6a2aed23fae46579cfb2d8ca1a87c8b7437f569aa6a8609d000b84ecce718ffd56fe32455293742c605e0eaa4c0f6180ba66713d13942d9ad53f249f856254eb957bea8129a9ab9a7bbf68b936861ec5ac3c38a040a5d71c8a1813bfc7be976f7ecd872463e40849c6da7efcab28b3d562caabf8bd053199747e3d3976d05aef111f7829a8858ff52c446db8f4a22699fe7deb27e9d5e88253d73ebff9eaf48d53cd6280f7e25c4464fe5527bf2a01893b29583dea1c01e127f49279d0004e2f8ff041d155fb863e3fa60f7d4fcce99c0f2250a8867ba5fcce207bc0f6ba4b461f000225dd5c9020d9e858e89fb6cedd7ca832d86b76f25d7dd000b19530742433762d7a7bb68f438f76aae65e5d6ecddd8075ce501c57f867296683580ab988c08e4becd50b1fe12d4a11bba9e49ca312ad7e008efbbaaa2e8befe6b4c1dba443bbe63e47e228ebab27a36f2ef102bb26b62173f2b20dd6b15d19a51bad824e0431a3a5a1b937e0c9ee460dc20c4313e1eacf9ddccdea94d82277a94c7d835c94354542aaa2f8439f0ba1a33e22db377ca6604888eb0652ca4ccfc22fc8dc6a29bbd1a915dfa5fd64357431b1150bbd66f4294a06d703c7952e0171188e74e55dd747d82a063e81711e74f31be9a17261d1f0a12b0174f4b199acd89a159f30dc376f2fe4d5c97a1bc0a20a98fd00c890f55367ec7594c13a1c9867f8d43c63a9c68ffbb0beca8e99e465eb4a5229473bf7393bbfcaa6f12095c923ef29fecdd349998b8d703c447f5167f73e27cecb3e3736e256e56439020abb577c1b06ade84f03c5cfc3358be2cd502a9bea11216bd653b6752ba51d4a4dc060eb1841d915df6259d78215e5183ea5732f31dc193a3533274fd3db347d4a223a2a58db07dbcedbc671ebf8d68ae88be87361df5405c22460836a9c092487facdad6ca36ced1a0c16ddd966465abcad5120deeafaaa5204ee359775a72971106b982c12e0b174b12bee0cbe8b518c4b922a04d9b9f1e938ff1ef3135da0a627b1b880bf8c1571bcfc68b104f35022e9236e868173cc5791121726e93b21cfacc8315565bd9f3f47164be4a10c3c301b39a45086f74c44e659523639909facdc0a6685c44f0724c5b29b3dfd1a78e87a700e48816a010d48ca154579c6c398834cd63c7eca1d0fae72cad31721f93bffb24ef94c0f60661bfbe2a9847f7444761be357abc54f7a1b63157cf78af7edf71c1513d18410f69ca5c6f83d42a073a4b3c58182b912e940d015feec12ee686f8dc8d55373a75532703e1200cae5a5a10353a7f4287a639065877c11f026b8c98f15c0f094776caf46082bece871ed7e527fe59a7fec0fbb0cd106440428c49cd86ba4638520db16118c8e966c950e51f6ec322a55af76f8149db9a42a2d8cfe1e648d35725580f311970b0f5cba0106d1835e484585c869937c31afae9eac1946573681bdb01b42814b7631f4a1916650c1ede632a81d99edf3a8d1ba67211aafc5ae9167482f3f4feff287b2d3cdddcb3cba92e1cb95257663040aac3800b2a61fd445394e2a826b741b12b1b28656072847d17a58cbf8523ec98d61407ef22e096322fb340511dff0cd4bb89da7cac6f02de9f25ba1e0c6584deeaee475508e7d288a2fe78785af710d6499deefb33cd8ee5b0d7265f63d4b5af9db2b65080397a4e8ba692be9e4bf54b34faf4b3d67f81cfd2506d2497d01fe5644b21084367a1ecb0491241a37cba7829ae18ea3d5a7cd676c1299a50484ee35d0cbaecaad07fc15b19bb74bcf312f5278415641813927b40b35d71910b810dd3d4de19ade8542dc9cd863622592f0e9c6931e87ea1923e260e4faca0ca98a75ae6d42c31febae6eb4f1254fa57f0bd6c1c94d96e0e5561817846c446b7ae5d252a2bba884d238c2b90bad56dfd98f00d9ac99d8516f32a8ac3de0bb6320bedd2c7568c864d7506d6cab2e2256cb33c25c58f7ffd99edfb1cb4d139cfdd26384d1797622de2d5c818e9580a4b4d93797503c012f31f205002749fd2a092e702023e53bd35f9804bca4eace49100a7124b414e9274637382e28ed514bfd47a82b93ce44008924baf8e8d765596940dc68c30b10e1e5af2fb709cd7ad4837b4c2917daae093c97261ec3c971672cc101503458b73f00ea95a8152c3a6602978d248f1ee2dc9440d359f3319e2c89313fbbc51304a7c7aa041115a94d84187c6f427d5c059dde1718d1c9751d33c400bc6268a6e85e09ce3f411a256e1fef5593061200eb38ffbf0f308b411d3563cda1b00de544956cce9252f9809ec5e11c0b8d248b80d44f88ab519d6e236813da47543c45234947408ae80bcd96dea9105436eb908d1c6ec198457c83379ed9af52d45134acbde74892fb4407837b931a6fc9f94141efd414a1aee02350c76bcb7812141b5f4b2825daa573bd7e7d3bf5c1f3d527f608939276b0ad69511aa2abd1477245e98fb8941b64f113cf2c80c809926bf3f90247239b77b20b2524caeeb8c60bb4fe992be005ec156a33aa145ccc656a57ff7a6944ad99051389ce5ac3752b474ed94343aceabc24233d428e8de81510e0f531b3772ce417d5665eb099d1ad4bc98d74d652d17a2fe0c1c32f1b60c519f80772c838c635f2c2e52efdf6b7454f970c8d6e4788fd129bae90d0e6fd691ec17c60245689c24c44d1e7f97b17e015c3ee24670351456ae805782a047a4b04aeda346b29f26764599ee5327546296edded870437654563d7359db0f3e1fc8fd610f7a4b3532e0727dc29547c906f305f0ebe8f2b721122a78c544cfa7a9dc18c63f5c05d8789ff1c5fa758fae0e19b395743e66589ba765e8a095f846d0cf9621d9544a7637d824d4a327620341795468c1cd3fd8083f5fc787dc5e68c8181e3c35e8dcf03f485e2ec636fea1b0621e4246aa84c420466da54b6aa263dc71c8a42619231deb95b5fc7da23aa2ae6c0a21e5bb0af90f1b68c8cf86d973849d186de2ff9c86bbe2a97e0628a816414c13d536a5a05e220aec9d230a81fbd8a7a00c722a580528c2e73a69ca1e3bfe83709e5697b5e930d1acb3def013f66c1a97d505952ade19f57deb6d8e8ed1ad753d1d36cbe77d7e96af110dd54ae1d519cee087b98d9fb2ae7efc41af328122070d222e2946021a04c93a60fded135db7d2bf3cffb97eafcb7e77d436fe4590ff924b317e2df0a393ef1df2b4aa9224591bb8ff2b717264daed3aa98372f4f708fb808952af1382fdeab89e069e2aa3ab7f5a19c957b14b5cc11f0a6d043c63ff3851e405fd9a837390b3208bb54752de242274d3b3ad7f5be08c028bd6cde14f79e8361896e7c763a6737d58ba5239555086b2c2907d2500a77ec77bccab968869c28aaa3dce3ddd199666b202b449fabdfb78a1f6dfd55952ec81df85a60b33496459c7683d9df74cf565ccb98b8b4bfb16581e506b13b9e7c69f49e28b0ac55c3560717c6ece6cff03bf92056e80d2ac55b124d22beb6b2b70cd21079b4c243a19cf156d274413e6436084264ee9e469d0b1edb809aeaf418a13471633a4570b6b841cf565bfd52d7947788f1438ef1a5fc38d71db5109972df1ff2d1496207dc1b793db99f2e2b4212f3f28b6f8e8b0c0ed5d6627efbd06a2b836c20202d780a4f605990aaca411da4d58d45d7160b0d3f87d17191a25b740c313df42d2a2e9a66df287de25d85ef7f071ff360b5228ee33ef4e8c4380905c9bfe2936c25a898a81c79835ef2ed827fa7bb60f03e4e8ba171fedf78938311556a2de53e390f10ec3e0ae1021191620ce0cafacbf7ab9953e430074d1ea7ade56d38388065d05cf99dfddbc477f292cde4e8fb1263e4cc2122556aaaa95b52b4e46fb3ad1e4ae898207740fa3670b8b95448a35dea1511360ea9d3e304e29a72c134cbcc321867cb1233ec1967ce7005a188c894aa871749279c99f57aab97077ad9e84768f5f3971dabd1aa21567faa42aeabfc83776d9f2af19740d33966dba09ad9c8e8653789dfdbc120bf5c0ed0491a271f6f7e8cef9cad346b96b44d3c24f2520c0babf17d657482ec532209672eeab7d4ae5775cbe7487f4ff185ced3f4d6c31ce48725dcf1218f0d01903a2e799fe385584610c0b2d827795a095a69c47ab2b8cb87ab88b2e8b3b1a6bdf7c426a6a00cbb824612d30ad764161c8036b78682ac69e7aaf26abea23030ae5bbf8715eb87fbc5edc87fedef24d41ebd8d5b33be475c19bfcf81112c8751193854782aef64cd7df4b217e566636fac1490769c67f5eddb5a215cdd9a37e1f69bced9e3e15d786c16e4c4b6c0a07fc645231ac6fe4365f5a2e51b539eff3026774caa6f352187958b25f87731d47b93b27f7669673ac6a3042cdcc19a957fe65cbae11643b5e49fc967cd39083d3c048b6ee1d5dc697949605427d2eafe843033cbbf1d47b99942253edcb39ec9065bd8461d5e2e7f6e7fa52d3208f606efe1b54ee642a2356f9d612d45d43b6147c8fa8ce1262dcd953b338116ff9f6d2a24d86647fdf173654da988cd5361344bdca68918fc0152a11ff6c891fc67306ff1d8d386998cd1c7e50fa4178b0822aba748df64549d2bee6542f20c9c95df5576c2e338d41ec52cacc00383f709758a419445a90dbc2e29322202d5b4e82c75b7eabf7678a74307e01611d0865a0b02ff365cad95218ec57a906d22af1a84873e9b85f8baf2629d186f8bafff034131f5e0f7d0d2d5faf18da74bf8d69ac18fb9bc83d98cd4c8ce329dbc3be10b8ba076208462dc33b072dc61d7ac160f0b69f34e1e26eb982348cb91aa00888bad93782dfe92505cded81ffd6a70c99423f138871031ea300b0f8c9924366a2b9b81d9ce4794e1ecadcbb8320d755f8dfa4ad4ea38f6461ce6fb0bcb5593a42c87125d5af57e090d656a4ed67f301bfe235c9f875efef4c0a6de8cb3cea9cbe5811042e97dccc6d111197a989baba2d0534145a3164a6746cbb36fa112343412b671ef723709497a227bb123eea1c577277510859ead71d9e90585002532ee19822b02fcfa08d82bfd4845dbd1ff35fedd52f3b05fa47a11be1b811b231a2ac4c119c0b228798071be1c191a8c94723efd3d72743ab329770c59f2e9c13908641469981b323356a8465c34cb628a70f01986a521e1d2a632c26e53b83d2cc4b0edecfc1e68c65383a75726c2d6c6973746c35313a687474703a2f2f6d6972726f72732e6b65726e656c2e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35373a687474703a2f2f6d6972726f722e6161726e65742e6564752e61752f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35333a687474703a2f2f6674702e69696e65742e6e65742e61752f7075622f617263686c696e75782f69736f2f323031312e30382e31392f36303a687474703a2f2f6d6972726f722e696e7465726e6f64652e6f6e2e6e65742f7075622f617263686c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6d6972726f722e6f707475732e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6674702e6279666c792e62792f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35333a687474703a2f2f617263686c696e75782e6d6972726f722e6b616e6761726f6f742e6e65742f69736f2f323031312e30382e31392f34353a687474703a2f2f617263686c696e75782e6333736c2e756670722e62722f69736f2f323031312e30382e31392f35393a687474703a2f2f6d6972726f722e6373636c75622e7577617465726c6f6f2e63612f617263686c696e75782f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e6974732e64616c2e63612f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6c6573732e636f6765636f2e6e65742f7075622f617263686c696e75782f69736f2f323031312e30382e31392f34383a687474703a2f2f6d6972726f72732e3136332e636f6d2f617263686c696e75782f69736f2f323031312e30382e31392f35313a687474703a2f2f6d6972726f722e626a74752e6564752e636e2f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f72362e626a74752e6564752e636e2f617263686c696e75782f69736f2f323031312e30382e31392f35343a687474703a2f2f7777772e6c617165652e756e616c2e6564752e636f2f617263686c696e75782f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e767073667265652e637a2f617263686c696e75782f69736f2f323031312e30382e31392f35313a687474703a2f2f6d6972726f72732e646f747372632e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6674702e65656e65742e65652f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f722e61636164656d6963612e66692f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f722e617263686c696e75782e66692f617263686c696e75782f69736f2f323031312e30382e31392f33393a687474703a2f2f6d69722e617263686c696e75782e66722f69736f2f323031312e30382e31392f37333a687474703a2f2f646973747269622d636f666665652e6970736c2e6a7573736965752e66722f7075622f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972312e617263686c696e75782e66722f617263686c696e75782f69736f2f323031312e30382e31392f34383a687474703a2f2f6d69726f69722e657a76616e2e66722f617263686c696e75782f69736f2f323031312e30382e31392f35383a687474703a2f2f617263686c696e75782e6d6972726f72732e6f76682e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f34343a687474703a2f2f617263686c696e75782e706f6c796d6f72662e66722f69736f2f323031312e30382e31392f34323a687474703a2f2f617263686c696e75782e6c696d756e2e6f72672f69736f2f323031312e30382e31392f34393a687474703a2f2f61727466696c65732e6f72672f617263686c696e75782e6f72672f69736f2f323031312e30382e31392f35353a687474703a2f2f6d6972726f722e64652e6c656173657765622e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f35313a687474703a2f2f6d6972726f722e6465766e7531312e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f35353a687474703a2f2f667470352e677764672e64652f7075622f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35393a687474703a2f2f6674702e68616c696661782e727774682d61616368656e2e64652f617263686c696e75782f69736f2f323031312e30382e31392f36353a687474703a2f2f6674702e686f73746575726f70652e64652f6d6972726f722f6674702e617263686c696e75782e6f72672f69736f2f323031312e30382e31392f36393a687474703a2f2f6674702d737475642e68732d6573736c696e67656e2e64652f7075622f4d6972726f72732f617263686c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6d6972726f72732e6e2d69782e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e73656c666e65742e64652f617263686c696e75782f69736f2f323031312e30382e31392f36383a687474703a2f2f6674702e73706c696e652e696e662e66752d6265726c696e2e64652f6d6972726f72732f617263686c696e75782f69736f2f323031312e30382e31392f36313a687474703a2f2f6674702e74752d6368656d6e69747a2e64652f7075622f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35363a687474703a2f2f6674702e756e692d6b6c2e64652f7075622f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f36303a687474703a2f2f6d6972726f722e6c6976696470656e6775696e2e636f6d2f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35343a687474703a2f2f617263686c696e75782e706f7369746976652d696e7465726e65742e636f6d2f69736f2f323031312e30382e31392f34383a687474703a2f2f617263686c696e75782e6d6972726f72732e756b322e6e65742f69736f2f323031312e30382e31392f36303a687474703a2f2f6674702e63632e756f632e67722f6d6972726f72732f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35343a687474703a2f2f6674702e6e7475612e67722f7075622f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6674702e6f74656e65742e67722f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35343a687474703a2f2f6d6972726f722e6373652e6969746b2e61632e696e2f617263686c696e75782f69736f2f323031312e30382e31392f36323a687474703a2f2f6674702e6865616e65742e69652f6d6972726f72732f6674702e617263686c696e75782e6f72672f69736f2f323031312e30382e31392f35353a687474703a2f2f6d6972726f722e69736f632e6f72672e696c2f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35343a687474703a2f2f6d6972726f72732e70726f6d65746575732e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f35383a687474703a2f2f6674702e6a616973742e61632e6a702f7075622f4c696e75782f417263684c696e75782f69736f2f323031312e30382e31392f33353a687474703a2f2f617263686c696e75782e6b7a2f69736f2f323031312e30382e31392f35313a687474703a2f2f6d6972726f722e796f6e67626f6b2e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f34343a687474703a2f2f617263686c696e75782e676f6f64736f66742e6c762f69736f2f323031312e30382e31392f34373a687474703a2f2f617263686c696e75782e6d6972726f722e726f6f742e6c752f69736f2f323031312e30382e31392f34383a687474703a2f2f6d6972726f722e69686f73742e6d642f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f7273332e6b65726e656c2e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35353a687474703a2f2f6d6972726f722e6e6c2e6c656173657765622e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f36343a687474703a2f2f6674702e6e6c7575672e6e6c2f7075622f6f732f4c696e75782f64697374722f617263686c696e75782f69736f2f323031312e30382e31392f35333a687474703a2f2f617263686c696e75782e6e617574696c652e6e632f617263686c696e75782f69736f2f323031312e30382e31392f34323a687474703a2f2f6d6972726f722e617263686c696e75782e6e6f2f69736f2f323031312e30382e31392f36383a687474703a2f2f70696f74726b6f736f66742e6e65742f7075622f6d6972726f72732f6674702e617263686c696e75782e6f72672f69736f2f323031312e30382e31392f35353a687474703a2f2f6674702e726e6c2e6973742e75746c2e70742f7075622f617263686c696e75782f69736f2f323031312e30382e31392f35363a687474703a2f2f6d6972726f72732e61646e657474656c65636f6d2e726f2f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f722e617263686c696e75782e726f2f617263686c696e75782f69736f2f323031312e30382e31392f35383a687474703a2f2f6674702e726f6564752e6e65742f6d6972726f72732f617263686c696e75782e6f72672f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e776f726c6469732e6d652f617263686c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6d6972726f722e79616e6465782e72752f617263686c696e75782f69736f2f323031312e30382e31392f34383a687474703a2f2f6674702e6b616973742e61632e6b722f417263684c696e75782f69736f2f323031312e30382e31392f35383a687474703a2f2f73756e736974652e726564697269732e65732f6d6972726f722f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f7273342e6b65726e656c2e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35353a687474703a2f2f6674702e6c797361746f722e6c69752e73652f7075622f617263686c696e75782f69736f2f323031312e30382e31392f36323a687474703a2f2f6674702e706f72746c616e652e636f6d2f7075622f6f732f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f34323a687474703a2f2f617263686c696e75782e70757a7a6c652e63682f69736f2f323031312e30382e31392f35333a687474703a2f2f6c696e75782e63732e6e6374752e6564752e74772f617263686c696e75782f69736f2f323031312e30382e31392f35353a687474703a2f2f736861646f772e696e642e6e746f752e6564752e74772f617263686c696e75782f69736f2f323031312e30382e31392f35333a687474703a2f2f6674702e746b752e6564752e74772f4c696e75782f417263684c696e75782f69736f2f323031312e30382e31392f34393a687474703a2f2f6674702e6c696e75782e6f72672e74722f617263686c696e75782f69736f2f323031312e30382e31392f36303a687474703a2f2f6674702e6c696e75782e6b6965762e75612f7075622f4c696e75782f417263684c696e75782f69736f2f323031312e30382e31392f34333a687474703a2f2f617263686c696e75782e7375707365632e6f72672f69736f2f323031312e30382e31392f34393a687474703a2f2f63616b652e6c69622e6669742e6564752f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f72732e6361742e7064782e6564752f617263686c696e75782f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e6563652e76742e6564752f617263686c696e75782f69736f2f323031312e30382e31392f35373a687474703a2f2f7777772e67746c69622e6761746563682e6564752f7075622f617263686c696e75782f69736f2f323031312e30382e31392f36313a687474703a2f2f6d6972726f722e616e636c2e6861776169692e6564752f6c696e75782f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f7273312e6b65726e656c2e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f7273322e6b65726e656c2e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6c75672e6d74752e6564752f617263686c696e75782f66747066756c6c2f69736f2f323031312e30382e31392f35303a687474703a2f2f6d6972726f722e6d6f636b65722e6f72672f617263686c696e75782f69736f2f323031312e30382e31392f35313a687474703a2f2f6674702e6f73756f736c2e6f72672f7075622f617263686c696e75782f69736f2f323031312e30382e31392f34373a687474703a2f2f6d6972726f722e7269742e6564752f617263686c696e75782f69736f2f323031312e30382e31392f35323a687474703a2f2f6d6972726f72732e727574676572732e6564752f617263686c696e75782f69736f2f323031312e30382e31392f35353a687474703a2f2f6d6972726f72732e6c6178312e74686567636c6f75642e636f6d2f617263682f2f69736f2f323031312e30382e31392f35353a687474703a2f2f6d6972726f722e79656c6c6f7766696265722e6e65742f617263686c696e75782f69736f2f323031312e30382e31392f34363a687474703a2f2f6d6972726f72732e73742e757a2f617263686c696e75782f69736f2f323031312e30382e31392f6565"

func Test_Bencode(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)

    var data map[string]any

    td, err := hex.DecodeString(torrentData)
    if err != nil {
        t.Error(err.Error())
    }

    err = Unmarshal(td, &data)
    if err != nil {
        t.Error(err.Error())
    }

    test1 := "http://tracker.archlinux.org:6969/announce"
    assertEqual(data["announce"], test1, "RijndaelPKCS7Padding-res")

    test2 := "Arch Linux 2011.08.19 (www.archlinux.org)"
    assertEqual(data["comment"], test2, "RijndaelPKCS7Padding-res")

    test3 := "mktorrent 1.0"
    assertEqual(data["created by"], test3, "RijndaelPKCS7Padding-res")

    test31 := int64(1313750171)
    assertEqual(data["creation date"].(int64), test31, "RijndaelPKCS7Padding-res")

    test5 := "archlinux-2011.08.19-netinstall-i686.iso"
    assertEqual(data["info"].(map[string]any)["name"], test5, "RijndaelPKCS7Padding-res")
}

func Test_Data(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    var data Data

    td, err := hex.DecodeString(torrentData)
    if err != nil {
        t.Error(err.Error())
    }

    err = Unmarshal(td, &data)
    if err != nil {
        t.Error(err.Error())
    }

    test0 := []string{
        "announce",
        "comment",
        "created by",
        "creation date",
        "info",
        "url-list",
    }
    assertEqual(data.GetKeys(), test0, "Data-GetKeys")

    test1 := "http://tracker.archlinux.org:6969/announce"
    assertEqual(data.GetAnnounce(), test1, "Data-GetAnnounce")

    test2 := "Arch Linux 2011.08.19 (www.archlinux.org)"
    assertEqual(data.GetComment(), test2, "Data-GetComment")

    test3 := "mktorrent 1.0"
    assertEqual(data.GetCreatedBy(), test3, "Data-GetCreatedBy")

    test31 := int64(1313750171)
    assertEqual(data.GetCreationDate(), test31, "Data-GetCreationDate")

    test32 := "2011-08-19 18:36:11 +0800 CST"
    assertEqual(fmt.Sprintf("%s", data.GetCreationDateTime()), test32, "Data-GetCreationDateTime")

    test33 := []string{
        "length",
        "name",
        "piece length",
        "pieces",
    }
    assertEqual(data.GetInfoKeys(), test33, "Data-GetInfoKeys")

    test5 := "archlinux-2011.08.19-netinstall-i686.iso"
    assertEqual(data.GetInfoItem("name"), test5, "Data-GetInfoItem")

    assertNotEmpty(data.ToArray(), "Data-ToArray")
    assertNotEmpty(data.ToJSON(), "Data-ToJSON")
    assertNotEmpty(data.ToInfoArray(), "Data-ToInfoArray")
    assertNotEmpty(data.ToInfoJSON(), "Data-ToInfoJSON")
    assertNotEmpty(data.String(), "Data-String")

    data = data.SetItem("key111", "test-data")
    assertEqual(data.GetItem("key111"), "test-data", "Data-SetItem")
    data = data.SetAnnounce("test-SetAnnounce")
    assertEqual(data.GetItem("announce"), "test-SetAnnounce", "Data-SetAnnounce")
    data = data.SetComment("test-SetComment")
    assertEqual(data.GetItem("comment"), "test-SetComment", "Data-SetComment")
    data = data.SetCreatedBy("test-SetCreatedBy")
    assertEqual(data.GetItem("created by"), "test-SetCreatedBy", "Data-SetCreatedBy")

    test315 := int64(1313750175)
    data = data.SetCreationDate(test315)
    assertEqual(data.GetCreationDate(), test315, "Data-SetCreationDate")

    test316 := int64(1313750177)
    data = data.SetCreationDateTime(time.Unix(test316, 0))
    assertEqual(data.GetCreationDate(), test316, "Data-SetCreationDateTime")

    test317 := map[string]any{
        "test1": "11111111",
        "test2": "222222222",
    }
    data = data.SetInfo(test317)
    assertEqual(data.ToInfoArray(), test317, "Data-SetInfo")
}

func Test_SingleTorrent(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)

    var data SingleTorrent

    td, err := hex.DecodeString(torrentData)
    if err != nil {
        t.Error(err.Error())
    }

    err = Unmarshal(td, &data)
    if err != nil {
        t.Error(err.Error())
    }

    test1 := "http://tracker.archlinux.org:6969/announce"
    assertEqual(data.Announce, test1, "RijndaelPKCS7Padding-res")

    test2 := "Arch Linux 2011.08.19 (www.archlinux.org)"
    assertEqual(data.Comment, test2, "RijndaelPKCS7Padding-res")

    test3 := "mktorrent 1.0"
    assertEqual(data.CreatedBy, test3, "RijndaelPKCS7Padding-res")

    test31 := int64(1313750171)
    assertEqual(data.CreatDate, test31, "RijndaelPKCS7Padding-res")

    test5 := "archlinux-2011.08.19-netinstall-i686.iso"
    assertEqual(data.Info.Name, test5, "RijndaelPKCS7Padding-res")
}

func Test_MultipleTorrent(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)

    var data MultipleTorrent

    td, err := hex.DecodeString(torrentData)
    if err != nil {
        t.Error(err.Error())
    }

    err = Unmarshal(td, &data)
    if err != nil {
        t.Error(err.Error())
    }

    test1 := "http://tracker.archlinux.org:6969/announce"
    assertEqual(data.Announce, test1, "RijndaelPKCS7Padding-res")

    test2 := "Arch Linux 2011.08.19 (www.archlinux.org)"
    assertEqual(data.Comment, test2, "RijndaelPKCS7Padding-res")

    test3 := "mktorrent 1.0"
    assertEqual(data.CreatedBy, test3, "RijndaelPKCS7Padding-res")

    test31 := int64(1313750171)
    assertEqual(data.CreatDate, test31, "RijndaelPKCS7Padding-res")

    test5 := "archlinux-2011.08.19-netinstall-i686.iso"
    assertEqual(data.Info.Name, test5, "RijndaelPKCS7Padding-res")
}

func Test_NewEncoder(t *testing.T) {
    buf := bytes.NewBufferString("123")

    en := NewEncoder(buf)
    if en == nil {
        t.Error("NewEncoder fail")
    }
}

func Test_NewDecoder(t *testing.T) {
    buf := bytes.NewBufferString("123")

    en := NewDecoder(buf)
    if en == nil {
        t.Error("NewDecoder fail")
    }
}

func Test_Marshal(t *testing.T) {
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    buf := map[string]any{
        "abc": "test123",
        "efg": "test333",
    }

    res, err := Marshal(buf)

    assertError(err, "Marshal")
    assertNotEmpty(res, "Marshal")
}

func Test_MustMarshal(t *testing.T) {
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    buf := map[string]any{
        "abc": "test123",
        "efg": "test333",
    }

    res := MustMarshal(buf)

    assertNotEmpty(res, "MustMarshal")
}

func Test_SingleTorrent_Check(t *testing.T) {
    errChek := cryptobin_test.AssertErrorT(t)
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    now := time.Now()

    st := SingleTorrent{
        Announce: "Announce",
        AnnounceList: [][]string{
            {
                "AnnounceList-c1",
                "AnnounceList-c2",
            },
            {
                "AnnounceList-d1",
                "AnnounceList-d2",
            },
        },
        CreatDate: now.Unix(),
        Comment: "Comment",
        CreatedBy: "github deatil",
        Info: SingleInfo{
            Pieces: "PiecesklolkoijuikjPiecesklolkoijuikj1w32",
            PieceLength: 1,
            Length: 2,

            Name: "Name",
            NameUtf8: "Utf-8",

            Publisher: "Publisher",
            PublisherUtf8: "Utf-8",

            PublisherUrl: "PublisherUrl",
            PublisherUrlUtf8: "Utf-8",

            MD5Sum: "MD5Sum",
            Private: false,
        },
        Nodes: [][]any{
            {
                "qqqqqqqqq",
                "bbbbbbb",
            },
            {
                "yyyyy",
            },
        },
        Encoding: "Utf-8",
        CommentUtf8: "Utf-8",
    }

    test1 := []string{
        "AnnounceList-c1",
        "AnnounceList-c2",
        "AnnounceList-d1",
        "AnnounceList-d2",
    }
    assertEqual(st.GetAnnounceList(), test1, "Test_SingleTorrent_Check-GetAnnounceList")

    tlayout := "2006-01-02 15:04:05"

    assertEqual(st.GetCreationDateTime().Format(tlayout), now.Format(tlayout), "Test_SingleTorrent_Check-GetCreationDateTime")

    now2 := now.AddDate(0, 0, 1)
    st = st.SetCreationDateTime(now2)
    assertEqual(st.GetCreationDateTime().Format(tlayout), now2.Format(tlayout), "Test_SingleTorrent_Check-GetCreationDateTime-now2")

    hh, err := st.GetInfoHash()
    errChek(err, "Test_SingleTorrent_Check-GetInfoHash")
    assertNotEmpty(hh[:], "Test_SingleTorrent_Check-GetInfoHash")

    assertNotEmpty(st.GetInfoHashString(), "Test_SingleTorrent_Check-GetInfoHashString")

    pha, err := st.Info.GetPieceHashes()
    errChek(err, "Test_SingleTorrent_Check-GetPieceHashes")
    assertNotEmpty(pha, "Test_SingleTorrent_Check-GetPieceHashes")

}

func Test_MultipleTorrent_Check(t *testing.T) {
    errChek := cryptobin_test.AssertErrorT(t)
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    now := time.Now()

    st := MultipleTorrent{
        Announce: "Announce",
        AnnounceList: [][]string{
            {
                "AnnounceList-c1",
                "AnnounceList-c2",
            },
            {
                "AnnounceList-d1",
                "AnnounceList-d2",
            },
        },
        CreatDate: now.Unix(),
        Comment: "Comment",
        CreatedBy: "github deatil",
        Info: MultipleInfo{
            Pieces: "PiecesklolkoijuikjPiecesklolkoijuikj1w32",
            PieceLength: 1,
            Length: 2,

            Name: "Name",
            NameUtf8: "Utf-8",

            Files: []MultipleInfoFile{
                MultipleInfoFile{
                    Length: 22,
                    Path: []string{
                        "pp",
                    },
                    PathUtf8: "Utf-8",
                },
            },
        },
        Nodes: [][]any{
            {
                "qqqqqqqqq",
                "bbbbbbb",
            },
            {
                "yyyyy",
            },
        },
        Encoding: "Utf-8",
        CommentUtf8: "Utf-8",
    }

    test1 := []string{
        "AnnounceList-c1",
        "AnnounceList-c2",
        "AnnounceList-d1",
        "AnnounceList-d2",
    }
    assertEqual(st.GetAnnounceList(), test1, "Test_MultipleTorrent_Check-GetAnnounceList")

    tlayout := "2006-01-02 15:04:05"

    assertEqual(st.GetCreationDateTime().Format(tlayout), now.Format(tlayout), "Test_MultipleTorrent_Check-GetCreationDateTime")

    now2 := now.AddDate(0, 0, 1)
    st = st.SetCreationDateTime(now2)
    assertEqual(st.GetCreationDateTime().Format(tlayout), now2.Format(tlayout), "Test_MultipleTorrent_Check-GetCreationDateTime-now2")

    hh, err := st.GetInfoHash()
    errChek(err, "Test_MultipleTorrent_Check-GetInfoHash")
    assertNotEmpty(hh[:], "Test_MultipleTorrent_Check-GetInfoHash")

    assertNotEmpty(st.GetInfoHashString(), "Test_MultipleTorrent_Check-GetInfoHashString")

    pha, err := st.Info.GetPieceHashes()
    errChek(err, "Test_MultipleTorrent_Check-GetPieceHashes")
    assertNotEmpty(pha, "Test_MultipleTorrent_Check-GetPieceHashes")

    assertNotEmpty(st.Info.GetFileList(), "Test_MultipleTorrent_Check-GetFileList")

}
