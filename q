[33mcommit 421e1493d6be573a826ae3e79904e363478f3bb1[m[33m ([m[1;36mHEAD[m[33m -> [m[1;32mmain[m[33m)[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Nov 1 05:01:39 2024 +0100

     I’ll split this slice into five parts, take each part and pass it to the
    sum() function together with the channel c, and call it as a goroutine

[33mcommit eda1304199c81714b4e2188e174e1eb743a3fb15[m[33m ([m[1;31morigin/main[m[33m)[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Nov 1 02:28:52 2024 +0100

    Added goroutines to prevent deadlock

[33mcommit 802b3b98fed1c63408a7848794c0eb6f72caf28c[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Thu Oct 31 23:13:11 2024 +0100

    Tried creating a channel but a deadlock error was thrown, because I need a goroutine since If you try to send something into a channel, Go waits until something else is ready to receive it.
    In the original code, there was nothing ready to receive that 5 you sent, so Go got stuck waiting for a receiver.

[33mcommit 16051d48d87a62b0cb8ec7272b4b1503ed17f8ca[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Thu Oct 31 10:44:58 2024 +0100

    To print out the interim balance in the goroutines, we use the
    Mutex object to mark out the critical section as done earlier

[33mcommit 89f00a104e226ac5fd1869b2ac227f14a016fddd[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Wed Oct 30 15:09:46 2024 +0100

    Replaced Mutex locks with atomic counters to handle concurrent updates
    to shared variables. Using tomic.AddInt64() provides a more efficient,
    lightweight solution for simple increments/decrements, improving performance
    by reducing lock contention. This change ensures safe concurrent access
    while simplifying the code.

[33mcommit 7a7cef00a99dde78ba5bb7c444c654e91fa1c6fd[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Tue Oct 29 22:13:49 2024 +0100

    Learnt about using mutexes to better sychronize goroutines

[33mcommit 99031f19e6064e827bdfa2976bc77f17ef8d6bf5[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Tue Oct 29 22:04:28 2024 +0100

    deleted simple say function and worked with credit() and debit() functions are updating a shared variable balance

[33mcommit 6b469d8382c506272f7b6b319bbfef252f4c82cd[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Tue Oct 29 21:49:05 2024 +0100

    practiced goroutines with simple say function

[33mcommit f8b915c3dcabebe235a7f345908cd39221200d11[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 28 16:25:38 2024 +0100

    Created file to practice and learn threading and go routines

[33mcommit 1590db60fae2f174a763504ca0a2b620f4284a93[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 28 16:01:03 2024 +0100

    Determining whether a value implements

[33mcommit ce82d19a66871bc60b1f71c843e6aec96c9d29ca[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 28 12:29:15 2024 +0100

    ädd several struct and created the shape interface to host the Area() method then created a calculateArea function that takes in a slice of type []Shape(he shape interface, so it'll be easy to calculate areas of any other new shapes that might arise)

[33mcommit bd821ece0c0689ff42008841a6dbd8d6c67d66dc[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Sun Oct 27 21:45:17 2024 +0100

    Trying to better understand the stringer interface with chatgpt

[33mcommit 6c1439b22f42ca1ed54c56ab7aad1f5689626fa5[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Sat Oct 26 21:07:42 2024 +0100

    Implementing an interface

[33mcommit 47436b2ac21c256c9481d836e56b4797dfc95448[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 25 17:41:49 2024 +0100

    Understood how json.MarshallIndent is formatted

[33mcommit bbc33687e3feeb8b198f83a2d2b33344fdd148cd[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 25 12:32:22 2024 +0100

    Created interfaces.go to learn about creating method signatures with interfaces

[33mcommit 7b3f5f5303c6af8adc038ee114666042e7acaa02[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 25 12:14:21 2024 +0100

    Encoding to Json

[33mcommit f0221967630138fe7620447b697b07d9ebc99afb[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 25 12:07:27 2024 +0100

    Accessing a nested struct in json

[33mcommit c207af28dd68e7e03a3cb38ec38827d08504a336[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 25 00:25:23 2024 +0100

    Mapping custom attributes

[33mcommit d9258818b6e58975a64eba89e8501e075214dffb[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Thu Oct 24 15:12:38 2024 +0100

    created json.go to practice encoding and decoding data and started off with a straightforwardway to convert a Json string in GO using json.Unmarshal from the encoding/json package in GO

[33mcommit 50cb6a8df722b200a3af4ddf88dc39cd089d3a4c[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Thu Oct 24 12:15:35 2024 +0100

    Actualized sorting of map of structs with the new sliceMembers struct created

[33mcommit 7deb3eea23176927e4aaafba7e8b96549a8a955f[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Wed Oct 23 20:02:50 2024 +0100

    moving on to ORTING A MAP OF STRUCTS, also checking why git streaks aren't be added

[33mcommit 362fb68019439ebe1c4544c02542f688309b6cf4[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Wed Oct 23 19:44:18 2024 +0100

    Using structs and maps together

[33mcommit a888922011faa62cbfc86c051465654d4426d9be[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Tue Oct 22 16:36:27 2024 +0100

    implemented delete function and used a for range to iterate through heights map

[33mcommit d7a55a7a72c7be0c88b584e74695ad90e4ff9bba[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 22:10:25 2024 +0100

    Initializing Maps

[33mcommit 527023a45eb181ffa515612b8bc76ecd862cb040[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 21:44:53 2024 +0100

    Elaborated on methods, defined variable and pointer receivers

[33mcommit 3d60a73880e227a342772608f79419b63f140782[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 14:15:24 2024 +0100

    Imported cmp package to have more flexibility in comparing structs when fields aren't comparable, the cmp package allows me to override the Equal functionso that I can implement my own custom comparer of structs

[33mcommit 8defda2e5cb613c04191d97af0bbe912a4a0850c[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 01:51:19 2024 +0100

    Idiomatic Go Code, hence I need constructor functions to create and initialize structs

[33mcommit 8e40333b26e2451b87f0ef688ac0e1c7563e0952[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 01:46:13 2024 +0100

    Learning about structs, chapter 7, Go Programming Language for Dummies

[33mcommit 6596de6b833ba5d4fecf0c8fb180dbb22b5ab99d[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 01:01:23 2024 +0100

    Added delete function that deletes an element from a slice, and tested it against all possibles cases and ensured proper error management

[33mcommit 10d89620f8d44042d157edf4f0ceebff8b874b91[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Mon Oct 21 00:44:33 2024 +0100

    Implementing an insert functin that uses append() for inserting items into a slice

[33mcommit 9aecd7a624f84681d39a3603c19b791fd969c1c0[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Sun Oct 20 23:19:17 2024 +0100

    learning git from scratch again, created git_practice to track, Version 1

[33mcommit 6718a875c38ebfab27d9129e6cae9791ee51a840[m
Author: Benita Emudianughe <emusbeny@gmail.com>
Date:   Fri Oct 18 20:08:36 2024 +0100

    Learning Go, recently working with a textbook and the ztm youtube video
