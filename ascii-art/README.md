Project Description:
The ASCII Art program written in go language, is designed to take a string as an argument and output the string in a graphic representation using ASCII characters. This graphical representation is created by rendering the input string with ASCII characters to form a visual representation.

File Structure: main folder(asciiart), contains(3 banner files; shadow.txt,standard.txt and thinkertoy.txt, a ReadMe.md file containing the documantaion of the program, 5 function folders with their test files and a go mod file that contains modules that help in runing the program )

Features:


    1. Accepts strings containing numbers, letters, spaces, special characters, and newline characters (\n).
    2. This program only takes strings containing ASCII character.
    3. Supports 3 banner file (files containig various graphical representaions of Ascii characters)as shown in the:
        shadow.txt
        standard.txt
        thinkertoy.txt
    4. Test files are meant to test the functionality of the project to ensure efficient operation of the functions.
    5. The program uses the standard banner file as adefault banner file if the user does not specify the desired banner file.
    6. The program has five functions(codes that help achieve a specific functionality/task within a program) including : 
                1. augument ;checks if the correct arguments is passed in the commadline and if the string passed as an argument only contains ASCII cahaters, it also handles cases of newline.
                2. bannerfile ; gets the selected banner file from the commadline argument if non is specified it takes a default(standard.txt)
                3. printascii ; prints the ascii art as an output
                4. readfile ; reads the banner file
                5. splitlines ; splits the content of the selected banner file with a new line

Usage:
To run the program, follow these steps:

    1.Clone the repository by copying this to the terminal : git clone https://learn.zone01kisumu.ke/git/bshisia/ascii-art.git

    2.Navigate to the program folder by copying the command below to the terminal 

                    cd asciiart

    3.Run the program using the following command/ copy the command below to the terminal

                    go run . "desired string" specify bannerfile

    4. To display the output with newline characters visible, you can add | cat -e at the end of your last argument like shown below 

                    go run . "\n" | cat -e

Example:
For example, running the program with the commands below

command 1.

go run . "Hello Alice!" shadow

would output an ASCII representation of the string "Hello, ASCII Art!" using the shadow banner style as shown below

                                                                           
_|    _|          _| _|                  _|_|   _| _|                   _| 
_|    _|   _|_|   _| _|   _|_|         _|    _| _|      _|_|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|       _|_|_|_| _| _| _|       _|_|_|_| _| 
_|    _| _|       _| _| _|    _|       _|    _| _| _| _|       _|          
_|    _|   _|_|_| _| _|   _|_|         _|    _| _| _|   _|_|_|   _|_|_| _| 
                                                                           

command 2.

go run . "Hello Brian!" thinkertoy

would output an ASCII representation of the string "Hello, ASCII Art!" using the thinkertoy banner style as shown below
                                               
o  o     o o           o--o                  o 
|  |     | |           |   |     o           | 
O--O o-o | | o-o       O--o  o-o    oo  o-o  o 
|  | |-' | | | |       |   | |   | | |  |  |   
o  o o-o o o o-o       o--o  o   | o-o- o  o O 
   
command 3.

go run . "Hello Joe!" standard

would output an ASCII representation of the string "Hello, ASCII Art!" using the standard banner style as shown below

 _    _          _   _                      _                  _  
| |  | |        | | | |                    | |                | | 
| |__| |   ___  | | | |   ___              | |   ___     ___  | | 
|  __  |  / _ \ | | | |  / _ \         _   | |  / _ \   / _ \ | | 
| |  | | |  __/ | | | | | (_) |       | |__| | | (_) | |  __/ |_| 
|_|  |_|  \___| |_| |_|  \___/         \____/   \___/   \___| (_) 
  
Contributions:
This project was made possible by the contributions from :
Alice Okingo, Brian Shisia and Joe