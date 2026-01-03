# **Product Requirements Document: Flip (Native macOS)**

## **1\. Overview**

Product Name: Flip  
Domain: nsFlip.com  
Platform: macOS (Exclusive)  
Language: Go (Golang)  
Library: github.com/caseymrm/menuet  
Objective: Create a native-feeling menu bar application that toggles the system-wide "Natural Scrolling" setting.

## **2\. Problem Statement**

The user alternates between a trackpad (requiring "Natural" scrolling) and a mouse (requiring "Standard" scrolling). Changing this setting manually via System Settings is inefficient and high-friction.

## **3\. Goals & Success Metrics**

* **Native UX:** The app must look and behave like a native Apple menu bar extra.  
* **Instant Feedback:** The menu bar icon must dynamically update to reflect the current scroll state.  
* **Efficiency:** Toggling must require only a single click (or a simple menu selection).  
* **Persistence:** The app must read the source of truth from macOS defaults on startup.

## **4\. User Stories**

1. **As a user**, I want a native menu bar icon that clearly indicates if my scrolling is currently "Natural" or "Standard."  
2. **As a user**, I want to click the icon to see a native macOS dropdown menu.  
3. **As a user**, I want to select a toggle option (or click the icon directly if configured) to switch modes.  
4. **As a user**, I want the app to handle the "defaults" modification silently in the background.

## **5\. Functional Requirements**

### **5.1 The Interface (UI via Menuet)**

* **Menu Bar Icon:**  
  * **Dynamic State:** The icon must change based on the system state.  
    * *Icon A:* "Natural" (e.g., fingers swiping up).  
    * *Icon B:* "Standard" (e.g., mouse wheel).  
  * **Text Label (Optional):** Option to display text next to the icon (e.g., "Nat" or "Std") for clarity, leveraging menuet's text support.  
* **Dropdown Menu:**  
  * **Header:** Display "Scroll Direction" (grayed out title).  
  * **Option 1:** "Natural" (Checkmark visible if active).  
  * **Option 2:** "Standard" (Checkmark visible if active).  
  * **Divider:** Standard macOS separator.  
  * **Option 3:** "Quit Flip".

### **5.2 The Logic (Backend)**

* **Startup Routine:**  
  1. Initialize menuet.  
  2. Run defaults read NSGlobalDomain com.apple.swipescrolldirection.  
  3. Update the menuet state to match the result.  
* **Toggle Logic:**  
  1. User selects the alternative mode.  
  2. App runs defaults write ....  
  3. **Crucial Step:** App forces a preferences reload (system notification).  
  4. App updates the menuet checkmarks and main icon.

## **6\. Technical Implementation Details**

### **6.1 Tech Stack**

* **Language:** Go (requires CGO).  
* **GUI Library:** github.com/caseymrm/menuet  
  * *Why:* Provides direct access to NSMenu, NSMenuItem, and correct macOS event loops.  
* **Build Requirements:**  
  * Must be compiled on macOS.  
  * CGO\_ENABLED=1 is required.

### **6.2 System Interaction (Shell Commands)**

The core logic relies on interacting with the defaults system.

**A. Read Command:**

Bash

defaults read NSGlobalDomain com.apple.swipescrolldirection

**B. Write Command:**

Bash

defaults write NSGlobalDomain com.apple.swipescrolldirection \-bool \<true|false\>

C. Sync Command (The "CFPrefs" update):  
Because menuet is native, we might get away with just writing the default. However, to ensure other apps see the change, we should run:

Bash

/System/Library/PrivateFrameworks/SystemAdministration.framework/Resources/activateSettings \-u

*Note:* If that specific private framework command fails on newer macOS versions (Sequoia/Sonoma), we fall back to killall cfprefsd.

### **6.3 Menuet Architecture**

menuet uses a function that returns the menu structure. The app will look roughly like this:

Go

func menuItems() \[\]menuet.MenuItem {  
    return \[\]menuet.MenuItem{  
        {  
            Text: "Natural Scrolling",  
            Clicked: func() { setScroll(true) },  
            State: isNaturalState(), // Returns true if active  
        },  
        {  
            Text: "Standard Scrolling",  
            Clicked: func() { setScroll(false) },  
            State: \!isNaturalState(),  
        },  
        { Type: menuet.Separator },  
        { Text: "Quit", Clicked: menuet.Quit },  
    }  
}

## **7\. Risks & Constraints**

* **CGO Complexity:** Compiling with CGO can sometimes make cross-compilation harder (not an issue here since you are building on Mac for Mac).  
* **Icon Assets:** menuet works best with template images (PDF or PNG) that automatically adjust for Light/Dark mode. We will need to find or create simple SVG/PNG icons that adhere to Apple's interface guidelines.  
* **Latency:** There is a slight delay (milliseconds) between clicking the menu and the shell command executing. We should update the UI *optimistically* (immediately) rather than waiting for the shell command to return.

## **8\. Development Roadmap**

### **Phase 1: Logic Verification**

* Create a Go function GetScrollState() and SetScrollState(bool).  
* Ensure the shell commands actually toggle the behavior on your specific OS version.

### **Phase 2: Menuet Implementation**

* Set up the menuet.App().  
* Implement the dynamic checkmarks in the dropdown menu.  
* Implement the dynamic Menu Bar icon (changing the top-level icon based on state).

### **Phase 3: Polish**

* Add a "Launch on Login" feature (common requirement for utility apps).  
* Package into a .app bundle structure so it doesn't look like a terminal executable.

### ---

**Would you like me to scaffold the main.go file using menuet so you can start filling in the logic?**