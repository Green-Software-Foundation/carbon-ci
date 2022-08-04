# importing sys
import sys

"""
import dashboard
from dashboard import buffer

buffer.seek(0)
print(buffer.read())
"""

# adding Folder_2 to the system path
#sys.path.insert(0, 'C:\\Users\\lydia.catterall\\OneDrive - Avanade\\playground\\Carbon_CI_Pipeline_Tooling\\src\\carbon-measure-action\\carbon_measure.go')

#from carbon_measure.go import odd_even, add

#file2 = open('carbon_measure.go', 'r')

#file2.close()

file = open('sci_scores.csv', 'a')

# Simulating a list of links:
#example_data = ['11', '0', '3']

buildNum = 11
avg = 7
total = 5

file.write(str(buildNum) + ',' + str(avg) + ',' + str(total) + "\n")  

file.close()