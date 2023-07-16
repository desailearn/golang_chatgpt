Using Commandline:
        
        Pre-Req: golang must be installed on you computer
        
        Step 1: git clone "https://github.com/desailearn/golang_chatgpt.git"
        
        Step 2: Generate your API Key in ChatCPT (if don't have one)
        
        Step 3: open .env file and update your API_KEY (Save)
        
        Step 4: cd <golang_chatgpt>
        
        Step 5: Enter Command => "go run main.go"
            You should see the message "How can I help yo ? ('quit' to end):"
        
        This is an interactive Chat GPT! 
        
        See the examples.txt file for sample examples

Using VSCode:

        Pre-Req: VSCode and golang must be installed on your computer
        
        Step 1: Open the terminal (mac or Windows)
        
        Step 2: Type Code . (To open VSCode)
        
        Step 3: Open new Terminal in VSCode
        
        Step 4: Inside VSCode Terminal type command ==> git clone "https://github.com/desailearn/golang_chatgpt.git"
        
        Step 2:  Generate your API Key in ChatCPT (if don't have one)
        
        Step 3: open .env file update you API_KEY (Save)
        
        Step 4: Inside VSCode Terminal type command ==> cd <golang_chatgpt> (If you are in current dir)
        
        Step 5: Inside VSCode Terminal type command ==> "go run main.go"
            You should see the message "How can I help yo ? ('quit' to end):"
        
        This is an interactive Chat GPT! 



        Here are some salesforce examples:


Here are some salesforce examples:

**  Use Case 1: Create APEX and Triggers  **
    Create a batch apex class with the below requirements.	
          Create a task for a case owner that’s in an open state.
          Assign a due date to the  task as one week from now
          Create an email auto-trigger as a daily reminder until the case is closed
          After completing attach the completed case to the  task to case id “XXXXXXX”
          Write unit test cases for triggers and classes.
       
**        Command line Input:  (Screenshot)   **        
                ![Input](Input_sf_Usecase_1.png)

**        Response Output:        **
                **        Response Output (Screenshot - 1 ):        **
                       ![Output 1](Output_sf_usecase_1a.png)
                **        Response Output (Screenshot - 2 ):        **
                       ![Output 2](Output_sf_usecase_1b.png)
                **        Response Output (Screenshot - 3 ):        ** 
                       ![Output 3](Output_sf_usecase_1c.png)
                **        Response Output (Screenshot - 4 ):        **
                       ![Output 4](Output_sf_usecase_1d.png)   


Use Case 2: Write a salesforce trigger with the  below requirements:
        Whenever an  account is deactivated trigger an external API end-point with the below payload 
            accountID 
            accountname
            dateDisabled
            array: impactedContacts
        Diabale all the associated contacts

Example Output from the command line
Input:

Output:
