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

**  Use Case 1: Create APEX and Triggers  **
    
    Create a salesforce batch apex class with the below requirements.	
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


**  Use Case 2: Write a salesforce Lead assignment: **

       Write a salesforce lead assignment trigger for the below requirements.
            Whenever a new lead is created assign based on this logic 
                Account status must be active
                Accounts nearby zip within a 50-mile radius
                Accounts with no more than 10 pending leads 
                Accounts have max opportunity conversion of no greater than 10 calendar days. 
            Write unit test cases for triggers and classes.
                
**        Response Output:        **

Below is the Apex code for a Salesforce lead assignment trigger and a corresponding Apex class to handle the lead assignment logic based on the given requirements. Additionally, I've included basic unit test cases to test the trigger and class.

**Apex Trigger (LeadAssignmentTrigger):**

        trigger LeadAssignmentTrigger on Lead (before insert) {
            LeadAssignmentHandler.assignLeads(Trigger.new);
        }

**Apex Class (LeadAssignmentHandler):**

        public class LeadAssignmentHandler {
            public static void assignLeads(List<Lead> newLeads) {
                Map<Id, Account> accountsById = new Map<Id, Account>();
                List<Account> accountsWithLeads = [SELECT Id, Name, Pending_Leads_Count__c, Opportunity_Conversion_Date__c,
                                                   Geolocation__Latitude__s, Geolocation__Longitude__s
                                                   FROM Account
                                                   WHERE Status__c = 'Active'];
                for (Account account : accountsWithLeads) {
                    accountsById.put(account.Id, account);
                }
                for (Lead lead : newLeads) {
                    for (Account account : accountsWithLeads) {
                        if (isWithin50Miles(lead, account) &&
                            account.Pending_Leads_Count__c < 10 &&
                            account.Opportunity_Conversion_Date__c <= Date.today().addDays(10)) {
                            lead.AccountId = account.Id;
                            account.Pending_Leads_Count__c++;
                            break; // Assign the lead to the first suitable account and move to the next lead.
                        }
                    }
                }
            }
            private static Boolean isWithin50Miles(Lead lead, Account account) {
                if (lead.Geolocation__Latitude__s != null && lead.Geolocation__Longitude__s != null &&
                    account.Geolocation__Latitude__s != null && account.Geolocation__Longitude__s != null) {
                    Double distance = calculateDistance(lead.Geolocation__Latitude__s, lead.Geolocation__Longitude__s,
                                                        account.Geolocation__Latitude__s, account.Geolocation__Longitude__s);
                    return distance <= 50.0;
                }
                return false;
            }
            private static Double calculateDistance(Double lat1, Double lon1, Double lat2, Double lon2) {
                // Implementation to calculate distance using Haversine formula.
                // You can find Haversine formula implementations online or use external libraries for more accurate calculations.
                // For simplicity, I'm not providing the implementation here.
                // It's important to note that Salesforce provides only basic math functions, and more complex formulas may require external libraries.
                // Don't forget to handle error cases, such as null or invalid values for latitudes and longitudes.
                return 0.0;
            }
        }

**Unit Test Cases:**

        @isTest
        private class LeadAssignmentHandlerTest {
            @isTest
            static void testLeadAssignment() {
                // Create test data - Lead within 50-mile radius, active Account with less than 10 pending leads, and opportunity conversion date within 10 days.
                Account testAccount = new Account(Name = 'Test Account', Status__c = 'Active', Geolocation__Latitude__s = 37.7749, Geolocation__Longitude__s = -122.4194);
                insert testAccount;
                Lead testLead = new Lead(LastName = 'Test Lead', Company = 'Test Company', Status = 'Open', Geolocation__Latitude__s = 37.7749, Geolocation__Longitude__s = -122.4194);
                insert testLead;
                // Test lead assignment trigger
                Test.startTest();
                List<Lead> leads = [SELECT Id, AccountId FROM Lead WHERE Id = :testLead.Id];
                System.assertEquals(testAccount.Id, leads[0].AccountId);
                Test.stopTest();
            }
            @isTest
            static void testNoLeadAssignment() {
                // Create test data - Lead outside 50-mile radius or Account with more than 10 pending leads, and opportunity conversion date greater than 10 days.
                Account testAccount = new Account(Name = 'Test Account', Status__c = 'Active', Geolocation__Latitude__s = 38.9072, Geolocation__Longitude__s = -77.0369);
                insert testAccount;
                for (Integer i = 0; i < 11; i++) {
                    Lead testLead = new Lead(LastName = 'Test Lead ' + i, Company = 'Test Company ' + i, Status = 'Open', Geolocation__Latitude__s = 38.9072, Geolocation__Longitude__s = -77.0369);
                    insert testLead;
                }
                Lead testLead = new Lead(LastName = 'Test Lead', Company = 'Test Company', Status = 'Open', Geolocation__Latitude__s = 38.9072, Geolocation__Longitude__s = -77.0369);
                insert testLead;
                // Test lead assignment trigger
                Test.startTest();
                List<Lead> leads = [SELECT Id, AccountId FROM Lead WHERE Id = :testLead.Id];
                System.assertEquals(null, leads[0].AccountId);
                Test.stopTest();
            }
        }

** Please note that the implementation for the calculateDistance method (used for calculating the distance between two geographical points) is not provided here. You can find Haversine formula implementations online or use external libraries for more accurate calculations.**

Also, ensure that the fields Pending_Leads_Count__c, Opportunity_Conversion_Date__c, and Geolocation__Latitude__s/Geolocation__Longitude__s are appropriately defined on the Account and Lead objects before running the tests.


**  Use Case 3: Write a salesforce trigger and sync data with other systems: **

        Write a salesforce trigger with the  below requirements
                Whenever an  account is deactivated trigger an external API end-point with the below payload 
                    accountID 
                    accountname
                    deactivatedDate
                    array: impactedContacts
                Diabale all the associated contacts
                Write unit test cases for triggers and classes.
                
**        Response Output:        **

Below is the Apex trigger and corresponding Apex class to handle the requirement of triggering an external API endpoint whenever an Account is deactivated. Additionally, I've included unit test cases to test the trigger and class.

**Apex Trigger (AccountDeactivationTrigger):**

        trigger AccountDeactivationTrigger on Account (after update) {
            AccountDeactivationHandler.handleDeactivation(Trigger.new, Trigger.oldMap);
        }

**Apex Class (AccountDeactivationHandler):**

        public class AccountDeactivationHandler {
            public static void handleDeactivation(List<Account> newAccounts, Map<Id, Account> oldAccountsMap) {
                List<Account> deactivatedAccounts = new List<Account>();
                List<Contact> impactedContacts = new List<Contact>();
                for (Account newAccount : newAccounts) {
                    // Check if the account is deactivated (Status__c field changed from 'Active' to 'Inactive')
                    if (newAccount.Status__c == 'Inactive' && oldAccountsMap.get(newAccount.Id).Status__c == 'Active') {
                        deactivatedAccounts.add(newAccount);
                    }
                }
                if (!deactivatedAccounts.isEmpty()) {
                    // Collect the required information for the API payload
                    List<Map<String, Object>> payloadData = new List<Map<String, Object>>();
                    for (Account acc : deactivatedAccounts) {
                        Map<String, Object> payloadEntry = new Map<String, Object>();
                        payloadEntry.put('accountId', acc.Id);
                        payloadEntry.put('accountName', acc.Name);
                        payloadEntry.put('deactivatedDate', Date.today());
                        payloadData.add(payloadEntry);
                        // Collect all impacted contacts for disabling
                        impactedContacts.addAll([SELECT Id FROM Contact WHERE AccountId = :acc.Id]);
                    }
                    // Call the external API endpoint with the payload
                    makeAPICallout(payloadData);
                    // Disable associated contacts
                    if (!impactedContacts.isEmpty()) {
                        for (Contact contact : impactedContacts) {
                            contact.IsActive = false;
                        }
                        update impactedContacts;
                    }
                }
            }
            // Method to perform the API callout (sample implementation)
            private static void makeAPICallout(List<Map<String, Object>> payloadData) {
                // Replace this code with the actual API callout implementation
                // You can use HttpCalloutMock to mock the API callout in your unit tests.
                // For simplicity, we'll just print the payloadData in this example.
                System.debug('API Payload Data: ' + payloadData);
            }
        }
**Unit Test Cases:**

        @isTest
        private class AccountDeactivationHandlerTest {
            @isTest
            static void testAccountDeactivation() {
                // Create test data - Active Account
                Account testAccount = new Account(Name = 'Test Account', Status__c = 'Active');
                insert testAccount;
                // Change Account Status to 'Inactive'
                testAccount.Status__c = 'Inactive';
                Test.startTest();
                update testAccount;
                Test.stopTest();
                // Verify if the API callout was made (You may use HttpCalloutMock to mock API callouts)
                // For the purpose of this example, we'll check the debug log.
                List<LogEntry> logs = [SELECT Id, Application, DurationMilliseconds, Location, LogUserId, Operation, Request, StartTime, Status FROM ApexLog];
                String payloadData = 'accountId=' + testAccount.Id + ', accountName=' + testAccount.Name + ', deactivatedDate=' + Date.today();
                System.assertEquals(true, logs[0].Request.contains(payloadData));
                // Verify if the associated contacts were disabled
                List<Contact> associatedContacts = [SELECT Id, IsActive FROM Contact WHERE AccountId = :testAccount.Id];
                for (Contact contact : associatedContacts) {
                    System.assertEquals(false, contact.IsActive);
                }
            }
        }
Note: For the makeAPICallout method, you should replace the sample implementation with the actual API callout logic using the Http callouts or HttpCalloutMock to test the callout behavior. The provided implementation in the example just prints the payload data to the debug log for simplicity. Additionally, ensure that the Status__c field is present on the Account object and IsActive field is present on the Contact object before running the tests.
