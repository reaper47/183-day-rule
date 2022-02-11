# 183-Day Rule Calculator
A simple way to calculate the 183-day rule

# Installation

Download and execute the [release binary](https://github.com/reaper47/183-day-rule/releases) for your OS.

# Usage

If you do not have past travel dates to take into consideration, then execute the program and follow the prompts to calculate the 183-day rule. Otherwise, follow the following steps.

1. Create a text file named "history.txt" next to the binary. Alternatively, you can download the example file of the same name under the release page.

2. Populate the file with your past travel dates. Each line consists of three space-separated parts: a start date, an end date, and a yes/no to whether you were abroad during this time. The start date of the next line must be the end date of the previous. Dates must be in the dd/mm/yyyy format.
 For example,

```
10/01/2021 01/05/2021 no
01/05/2021 27/07/2021 yes
27/07/2021 11/12/2021 no
11/12/2021 08/02/2022 yes
```

3. Execute the program and follow the prompts to calculate the 183-day rule.

# Feedback

If you have any feedback, please reach out to us at macpoule@gmail.com or open an issue on GitHub.

