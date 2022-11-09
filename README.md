# Introduction

## Pool
Adalah implementasi design patern bernama object pool pattern
Digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya

## Map
struct sync.Map
aman menggunakan concurrent mengunakan goroutine
aman dari race condition

## Cond
adalah implementasi locking berbasis kondisi

## Atomic
package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent. Dari pada memakai mutex untuk data primitive mending ini.

## Timer
- Adalah representasi satu kejadian
- Ketika waktu timer sudah expire, maka event akan dikirim kedalam channel

## Time Ticker
- Ticker adalah representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel