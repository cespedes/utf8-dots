# utf8-dots

utf-dots is a simple program that shows a UTF-8 text representing a monochrome image, using Braille characters.

## Usage

If invoked with any text as argument, it first renders that text using a 7x13 font, and then displays that image:

    $ utf8-dots Hello, world
    ⡄ ⢠     ⢤  ⠠⡄                       ⢤    ⡄
    ⡧⠤⢼⢠⣒⣒⡄ ⢸   ⡇ ⡔⠒⢢        ⡆⡀⡆⡔⠒⢢⠐⡔⠒⠄ ⢸ ⢠⠒⠢⡇
    ⠇ ⠸⠘⠤⠤⠂⠠⠼⠤ ⠤⠧⠄⠣⠤⠜ ⡰⠖     ⠣⠣⠃⠣⠤⠜ ⠇  ⠠⠼⠤⠘⠤⠔⠇

If invoked with a PNG image as the first argument, it displays that image:

    $ utf8-dots gopher.png
                            ⣀⣀⣤⣤⣴⣶⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣤⣤⣀⡀                       
                      ⢀⣀⣤⣶⣾⣿⣿⣿⣿⠿⠿⠛⠛⠛⠛⠉⠉⠁⠉⠉⠉⠉⠉⠉⠛⠋⠛⠛⠻⠿⢿⣿⣿⣿⣶⣦⣄⡀                  
            ⣀⣀⣀⣀⡀  ⢀⣤⣶⣿⣿⡿⠿⠛⠉⠉                  ⢀⣀⣠⣤⣤⣤⣄⣉⠙⠻⢿⣿⣿⣶⣄⣤⣴⣶⣶⣶⣶⣶⣤⡀       
        ⢀⣴⣶⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿⠟⢋⣡⣴⣶⣾⣿⣶⣿⣷⣶⣦⣄          ⣠⣴⣾⣿⣿⣿⠿⠿⠿⣿⣿⣿⣷⣤⡈⠙⢿⣿⣿⣿⠛⠛⠛⠛⠻⢿⣿⣷⡄     
       ⣰⣿⣿⠟⠋⠁  ⣨⣿⣿⡿⠋⣠⣾⣿⣿⠿⠛⠛⠉⠉⠙⠙⠻⢿⣿⣿⣦⡀     ⢠⣾⣿⡿⠛⠉      ⠈⠙⠻⣿⣿⣆ ⠈⠻⣿⣷⣤⣤⡀  ⠹⣿⣿⣆    
      ⣼⣿⣿⠃ ⣰⣶⣶⣶⣿⡿⠋⢀⣾⣿⡿⠋          ⠈⠻⣿⣷⡄   ⣰⣿⣿⠋            ⠘⣿⣿⣆  ⠙⣿⣿⣿⣿⠄  ⠹⣿⣿⡀   
      ⣿⣿⠇ ⢸⣿⣿⣿⣿⡟⠁ ⣾⣿⡟⢀⣤⣤⣤⡀         ⢹⣿⣷  ⢠⣿⣿⢃⣴⣿⣿⣷⣆         ⠸⣿⣿   ⠈⢿⣿⣿⠁  ⢠⣿⣿⠃   
      ⣿⣿⣇ ⠈⢿⣿⣿⡟  ⢰⣿⣿⢰⣿⣿⣿⣿⣿⣆         ⣿⣿⡇ ⢸⣿⡟⣼⣿⣿⣿⣿⣿⡇         ⣿⣿⡇   ⠈⣿⣿⣇ ⣀⣾⣿⡟    
      ⠘⢿⣿⣦⣄⣸⣿⡿   ⠸⣿⣿⢸⣿⣿⣿⣿⣿⡟         ⣿⣿⠇ ⠸⣿⣿⡹⣿⣿⣿⣿⣿⠃        ⢠⣿⣿     ⠸⣿⣿⣾⣿⣿⠏     
       ⠈⠛⢿⣿⣿⣿⠃    ⢿⣿⣧⠻⢿⣿⡿⠟         ⣼⣿⡿   ⢻⣿⣷⡌⠛⠛⠛⠁        ⣠⣿⣿⠏      ⢻⣿⣿⠋       
          ⣼⣿⡟     ⠈⢿⣿⣷⣄         ⢀⣠⣾⣿⣟⣡⣴⣶⣶⣶⣿⣿⣿⣶⣤⣀     ⢀⣀⣤⣾⣿⡿⠋       ⠸⣿⣿⡀       
          ⣿⣿⡇       ⠙⢿⣿⣿⣶⣤⣤⣤⣤⣤⣴⣶⣿⣿⢿⣫⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⡿⠟⠋          ⣿⣿⡇       
         ⢠⣿⣿⠁         ⠈⠙⠛⠿⠿⠿⠿⠿⠟⠛⢋⣽⣿⣿⢿⣿⣿⣿⣿⣿⣿⡿⢿⣿⣷⣮⡉⠉⠉⠉⠉⠉              ⢹⣿⣷       
         ⢸⣿⣿                    ⢼⣿⡟⠁ ⠈⠉⠉⠉⠉⠁  ⠈⢻⣿⣷                   ⢸⣿⣿       
         ⢸⣿⣿                    ⢻⣿⣷⣤⣤⣶⣶⣿⣿⣿⣶⣦⣤⣀⣼⣿⡿                    ⣿⣿⡆      
         ⢸⣿⣿                     ⠙⢿⣿⣿⠟⠛⣿⣿⡏⠛⣿⣿⣿⠿⠟⠁                    ⣿⣿⡃      
         ⢸⣿⣿                      ⠸⣿⣿⡄ ⣿⣿⡇ ⣿⣿⡇                       ⣿⣿⡇      
         ⠸⣿⣿                       ⢻⣿⣷⣾⣿⣿⣷⣶⣿⣿⠃                       ⣿⣿⡇      
          ⣿⣿⡆                       ⠉⠛⠛⠉⠉⠛⠛⠋⠁                        ⣿⣿⡇      
          ⣿⣿⡇                                                        ⣿⣿⡇      
          ⣿⣿⡇                                                        ⣿⣿⡇      
          ⢿⣿⡇                                                        ⣿⣿⡇      
          ⢸⣿⣿                                                        ⣿⣿⡇      
          ⢸⣿⣿                                                        ⣿⣿⡇      
      ⣀⣤⣶⣶⣾⣿⣿                                                        ⣿⣿⣿⣿⣷⣦⣄  
    ⣠⣾⣿⡿⠿⠛⢻⣿⣿                                                        ⢻⣿⣏⠙⠛⢿⣿⣷⡄
    ⣿⣿⣿⣶ ⢀⣸⣿⣿                                                        ⢸⣿⣿⣀⢰⣿⣿⣿⡇
    ⠹⣿⣿⣿⣾⣿⣿⣿⣿                                                        ⢸⣿⣿⣿⣿⣿⣿⡿⠁
     ⠉⠛⠛⠛⠉⢸⣿⡿                                                        ⢸⣿⣿⠈⠉⠉⠉  
          ⣾⣿⡇                                                         ⣿⣿⡆     
          ⣿⣿⡇                                                         ⣿⣿⡇     
          ⣿⣿⡇                                                         ⣿⣿⡇     
          ⣿⣿⡇                                                         ⣿⣿⡇     
          ⣿⣿                                                          ⣿⣿⡇     
         ⢸⣿⣿                                                          ⣿⣿⡇     
         ⢸⣿⣿                                                         ⢀⣿⣿⠃     
         ⠘⣿⣿                                                         ⢸⣿⣿      
          ⣿⣿⡇                                                        ⣼⣿⡟      
          ⣿⣿⡇                                                        ⣿⣿⡇      
          ⢹⣿⣷                                                       ⣸⣿⣿       
          ⠘⣿⣿⡆                                                     ⢠⣿⣿⠃       
           ⢹⣿⣿⡀                                                   ⢀⣾⣿⡏        
            ⢻⣿⣷⡀                                                 ⢀⣾⣿⡟         
             ⢻⣿⣿⣄                                               ⣠⣿⣿⠏          
              ⠙⢿⣿⣷⣄                                           ⣠⣾⣿⡿⠃           
              ⢀⣴⣿⣿⣿⣷⣤⡀                                     ⣀⣴⣾⣿⡿⣿⣿⣦⡀          
            ⢀⣴⣿⣿⠟⠉⠙⠿⣿⣿⣷⣦⣤⡀                             ⣀⣤⣶⣿⣿⣿⠟⠉ ⣈⣻⣿⣿⡄         
            ⣾⣿⣿⣿⣿  ⢠⣴⣿⣿⣿⣿⣿⣿⣶⣶⣤⣄⣀⣀⢀            ⢀⣀⣀⣀⣤⣤⣶⣾⣿⣿⣿⠿⠛⠻⣿⣿⣦ ⢿⣿⣿⣿⣿         
            ⢿⣿⣿⣿⣃⣠⣾⣿⡿⠟⠉  ⠉⠙⠛⠻⠿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⡿⠿⠛⠛⠉⠉    ⠙⢿⣿⣷⣮⣿⣿⣿⠟         
            ⠈⠛⢿⣿⣿⣿⠿⠋            ⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉              ⠘⠻⠿⠿⠟⠁          