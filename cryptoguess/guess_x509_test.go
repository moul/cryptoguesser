package cryptoguess

import "fmt"

func ExampleX509PKIXPublicKey_Valid() {
	experiment := NewX509PKIXPublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
-----END PUBLIC KEY-----`))
	experiment.Run()
	fmt.Println(experiment.String())
	// Output: x509: DER encoded public key: *rsa.PublicKey: &{608306305637925796091751781912442202795683837080520039301522018599439471659350929129047641600014219325768030296625424424345448069328375537097656264726872409669346964631973169088061068456644801547206592340239258807863914253613496344410768210694037962571766118343861315323684256007940397868833247575637775332458247093703506688405094111650106757090455436430116545278572444268308518814399325068421795179779154776544483060305921707389366982780625572375062677808785374136435412820035834031972253361392848622755078064481496629846759926136289522486108922454312015140092203833352973278691362374496694051958883486816981498108038162601522808412224573435547318218332467974418426290024000729037299062720016251706252245691191531968379259424268761139372144924021186347356192038449842702607942990619778520030774483248942866999156530759187345407296334294992294697537874595373096228689277826995922909616024893893083086622170950733804265273320908045376212718531890144840341838709366205409245406814550340900430806361508644473289694271880520086363787049600744919598020494076506686596308552173179749407600927060687430580477162742989412335721907576120509336456000464431105614240090531626536958109202417157669870271270127509090280837018845155202819024829523 65537}
}

func ExampleX509PKIXPublicKey_Invalid() {
	experiment := NewX509PKIXPublicKey([]byte(`lorem ipsum`))
	experiment.Run()
	fmt.Println(experiment.String())
	// Output: x509: DER encoded public key: err: no PEM formatted block found
}
