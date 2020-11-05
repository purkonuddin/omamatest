# omamatest
# soal 0 - Polindrom
input dua bilangan yang dipisahkan dengan karakter spasi (a n) sehingga menghasilkan deret angka antara a dan n.
kemudian dari deret angka tersebut akan diseleksi bilangan yang polindrom,
outputnya akan  menghasilkan jumlah bilangan yang polindron 

exp:.
    Input:
    22 25

    Output:
    1

# soal 1 - sorting kode buku berdasarkan kategori buku
input deret kode buku secara acak yang dipisahkan dengan spasi, kode buku terdiri dari kode kategori, Judul buku, dan tinggi buku
kode buku: "9Z99"
    9 Kode kategori
    Z Judul Buku
    99 tinggi buku

kemudian sistem akan mengurutkan data tersebut berdasarkan urutan kategori,
kemudan diurutkan berdasarkan tnggi buku tertinggi ke rendah

exp:.
    input: 3A13 5X19 9Y20 2c18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14
    output: 0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2c18 5X19 5G14 3N21 3N20 3N14 3A13

# soal 2 - bilangan yang hilang dalam deret angka
iputkan deret angka bertipe array string, dimana salah satu blangan nya dihilangkan
sistem akan menghitung jumlah deret angka dengan rumus Sn = 1/2 * n * (a + un) kemudian membandingkannya dengan jumlah deret angka yang ada
(ui+u2+u3+...)
sehingga dihasilkan bilangan yang hilang

exp:. 
    deretAngka := []int{23, 24, 25, 26, 27, 28, 30}

    result:
        Input:  23242526272830
        Nilai yang hilang:  29

# soal 3 - webservice untuk soal no 2
jalankan aplikasi dengan perintah go run main.go kemudian enter,
buka postman atau browser dan ketikan http://localhost:8081
end point:
    get http://localhost:8081/index

# soal 4 - docker service
tools yang digunakan:
    docker toolbox
    postman atau browser
    windows powershell
    Visual studio code


buat file dengan nama Dockerfile,
ubah environment golang GOOS=windows menjadi GOOS=linux melalui windows powershell dengan perintah $Env:GOOS="linux", untuk cek perubahannya gunakan perintah go env.
compile file main.go dengan perintah go build main.go melalui windows powershell atau terminal text editor Viscode sehingga menghasilkan file baru dengan nama "main" tanpa extensi
buka Docker Quickstart Terminal, dan lakukan perintah untuk membuat docker images menggunakan Dockerfile:
    docker build -t <nama docker images>    -> untuk membuat docker image
    docker images           -> untuk melihat docker images yang telah dibuat
    docker run -it -p 8080:8081 <nama docker images>
    docker stop <nama docker images> -> untuk menghentikan 

untuk mengakses endpoint pada browser atau postman gunakan ip docker 192.168.100 dan port yang sudah di deklarasikan yaitu 8080

# test
command:
    go test

output:
if output_value same as right_falue(
    PASS
    ok      omamatest/<folder_name>     0.000s
)
else(
    --- FAIL: <Test_function_name> (0.00s)
    <file_name>.go:11: Expected 1.5, got  <output_value>
    FAIL
    exit status 1
    FAIL    omamatest/<folder_name>     0.036s
)