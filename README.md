-2A
1. Program ini berfungsi untuk menukar tiga input string yang dimasukkan pengguna
2. Program ini berfungsi untuk menunjukkan bahwa suatu tahun inputan user termasuk tahun kabisat atau bukan
3. Program ini berfungsi untuk menghitung volume dan luas permukaan sebuah bola berdasarkan input jari jari (radius) yang diberikan oleh user 
4. Program ini berfungsi untuk mengubah suhu celsius inputan dari user menjadi reamur, kelvin, dan fahreinheit
5. Program ini berfungsi untuk mengubah 2 baris inputan user, 1 baris pertama berisi 5 data integer yang akan dicetak dalam format karakter, baris kedua berisi 3 data char yang akan dicetak +1 setelah masing-masing char tersebut

-2B
1. Program ini berfungsi untuk membuktikkan 4 urutan warna inputan user sudah sesuai dengan yang seharusnya, jika sudah akan muncul boolean true dan sebaliknya
2. Program ini berfungsi untuk mencetak sebanyak n yang diinputkan user, nantinya user diminta untuk memasukkan nama bunga sebanyak n juga, yang kemudian akan dirangkai dengan "-" sesuai yang sudah diinputkan
3. Program ini berfungsi untuk mengecek apakah beban belanjaan yang dibawa oleh Pak Andi akan membuat motor oleng atau tidak dengan cara menentukan total beban minimal dan selisih dari beban kanan dan kiri
4. Program ini berfungsi untuk nilai akar 2 dari suatu bilangan inputan user

-2C
1. Program ini berfungsi untuk menghitung biaya pengiriman Pertama-tama program meminta user untuk menginputkan berat parcel dalam gram. Setelah itu berat parsel dikonversi ke kilogram dan gram sisa dengan cara membagi berat total dengan 1000, biaya dasar dihitung berdasarkan berat kilogram, di mana setiap kilogram dikenakan biaya Rp. 10.000
2. Program ini berfungsi untuk menghitung nilai akhir mata kuliah berdasarkan inputan dari user dimana program ini akan menampilkan nilai akhir berdasarkan abjad "A,AB,B,BC,C,C,E"
  Soal & Jawaban 
    1.	Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal? 
Jawaban: Program eror, dikarenakan program tersebut tidak konsisten, terutama pada variablel nam dideklarasikan sebagai float64, namun di dalam blok if, program mencoba mengassign string (seperti "A", "AB", dll.) ke nam.  
 
    2.	Apa saja kesalahan dari program tersebut? Mengapa demikian? 
Jelaskan alur program seharusnya! 
Jawaban:  
Kesalahan program:  
-	Program tidak konsisten, pada variablel nam dideklarasikan sebagai float64, namun di dalam blok if, program mencoba memasukkan string (seperti "A", "AB", dll.) ke nam. 
-	Kondisi if adalah independen, sehingga jika nam lebih besar dari beberapa batas, beberapa blok if akan dieksekusi secara berurutan, dan nilai akhir nmk akan ditimpa beberapa kali. 
-	Sebagai contoh, jika nam = 80.1, maka semua kondisi berikut akan terpenuhi: 
nam > 80 → nmk = "A" ,nam > 72.5 → nmk = "AB" ,nam > 65 → nmk = "B" ,nam > 57.5 → nmk = "BC" ,nam > 50 → nmk = "C" ,nam > 40 → nmk = "D" 
Akhirnya, nilai nmk akan menjadi "D", yang jelas tidak sesuai dengan spesifikasi yang diharapkan. 
 
Perbaikan Program:  
-	Gantilah nam = "A" dengan nmk = "A", dan seterusnya untuk semua kondisi. 
-	Gunakan else if untuk memastikan hanya satu kondisi yang terpenuhi dan mencegah nilai nmk ditimpa oleh kondisi berikutnya. 
 
 
    3. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.
Jawaban:
Input: 93.5 -> Output: A.
Input: 70.6 -> Output: B. 
Input: 49.5 -> Output: D.

3. Program ini berfungsi untuk membuktikkan suatu bilangan inputan dari user termasuk bilangan prima atau bukan dengan cara mencari faktor-faktor nya


  
