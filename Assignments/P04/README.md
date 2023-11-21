# Program 4: Image Downloader
## Samuel Olatunde

## Description

This Go program is designed to concurrently download a set of images from given URLs and save them to disk. The program includes two versions of the downloader:

1. **Sequential Version:** Downloads and saves each image one after the other.
2. **Concurrent Version:** Downloads and saves images concurrently using goroutines.

The goal is to observe the benefits of concurrency for I/O-bound tasks, comparing the time taken to download images sequentially vs. concurrently.

### Image URLs

1. [Unsplash Image 1](https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640)
2. [Unsplash Image 2](https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640)
3. [Unsplash Image 3](https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640)
4. [Unsplash Image 4](https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640)
5. [Pexels Image 1](https://images.unsplash.com/photo-1590068561151-2aa0b87cda13?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D)
6. [Pexels Image 2](https://images.unsplash.com/photo-1610035974356-3e9f2c818347?q=80&w=1935&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D)
7. [StockSnap.io Image](https://cdn.stocksnap.io/img-thumbs/960w/people-man_2W0L5IENXQ.jpg)
8. [Pexels Image 3](https://images.pexels.com/photos/751696/pexels-photo-751696.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1)
9. [Pixabay Image](https://cdn.pixabay.com/photo/2017/05/11/12/35/girl-2304038_1280.jpg)

### Results:
**Sequential Logic took 1083710.6 ms.**  
**Concurrent Logic took 192.0054 ms.**

## Files

|   #   | File                  | Description                                              |
| :---: | --------------------- | -------------------------------------------------------- |
|   1   | [main.go](https://github.com/SamOlatunde/4143-PLC/blob/main/Assignments/P04/main.go)             | Main program file containing the image downloader logic. |

## Requirements

- The program reads a list of image URLs from a file.
- Two versions of the downloader are implemented: sequential and concurrent.
- Downloaded images are saved to disk with unique names.
- Proper error handling is implemented for failed downloads or any other issues.
- The program measures and prints out the time taken by each version to complete the downloads.
