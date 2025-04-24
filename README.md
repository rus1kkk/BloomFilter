# Bloom Filter CLI

[![Go Version](https://img.shields.io/github/go-mod/go-version/rus1kkk/BloomFilter)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/rus1kkk/BloomFilter)](https://github.com/rus1kkk/BloomFilter/releases)

## Project Overview (EN)

Implementation of a probabilistic Bloom Filter data structure in Go with command line interface.

## 📦 Installation

### Build from source
```bash
git clone https://github.com/rus1kkk/BloomFilter.git
cd BloomFilter
go build -o bloomfilter
```
### Install using go
```bash
go install github.com/rus1kkk/BloomFilter@latest
```
## 🚀 Usage
```bash
BloomFilter bloomfilter <size> <hash_count> <file> <word>
```
### Parameters:
size - Size of bit array (recommended 10-100x expected number of elements)

hash_count - Number of hash functions to use (typically 3-5)

file - Path to text file

word - Word to check

### Example:
```bash
bloomfilter 1000 3 words.txt hello
```
## 📝 Sample File
words.txt:
```
hello world
this is test  
probabilistic data structure
```
## 📊 How It Works
1. Creates Bloom Filter with specified parameters

2. Adds words from file to the filter

3. Checks if target word exists

4. Returns result:

 + "Element X PROBABLY in the set" - possible match

 + "Element X is DEFINITELY NOT in the set" - confirmed absence

## 🛠 Technical Details
 + Uses FNV-1a hash functions (fnv.New64())

 + Pure Go implementation (only Cobra dependency)

 + File reading with word-by-word scanning

 + Memory-efficient design

## Обзор проекта (RU)

Реализация вероятностной структуры данных Bloom Filter на языке Go с интерфейсом командной строки.

## 📦 Установка

### Установка из исходников
```bash
git clone https://github.com/rus1kkk/BloomFilter.git
cd BloomFilter
go build -o bloomfilter
```
### Установка через go install
```bash
go install github.com/rus1kkk/BloomFilter@latest
```
## 🚀 Использование
```bash
BloomFilter bloomfilter <размер> <количество_хешей> <файл> <слово>
```
### Параметры:

размер - размер битового массива (рекомендуется в 10-100 раз больше ожидаемого количества элементов)

количество_хешей - число используемых хеш-функций (обычно 3-5)

файл - путь к текстовому файлу

слово - слово для проверки

### Пример:
```
bloomfilter 1000 3 words.txt hello
```
## 📝 Пример файла
words.txt:
```
hello world
this is test
probabilistic data structure
```
## 📊 Как это работает
1. При запуске создается Bloom Filter с указанными параметрами

2. Слова из файла добавляются в фильтр

3. Проверяется наличие указанного слова

4. Выводится результат:

 + "Element X PROBABLY  in the set" - возможное наличие
 + "Element X is DEFINITLY NOT in the set" - точное отсутствие

## 🛠 Технические детали
 + Использует хеш-функции FNV-1a (fnv.New64())

 + Реализация на чистом Go без внешних зависимостей (кроме Cobra)

 + Чтение файла построчно с разбиением на слова

 + Оптимальное использование памяти


