local vimutil = require("vim.lsp.util")
local input = [[
frequent HEAD doc‧tor1 /ˈdɒktə $ ˈdɑːktər/  ●●● S1 W1 noun [countable]
                 
                 
Sense(doctor__1):1  (written abbreviation Dr) someone who is trained to treat people who are ill → GP
EXAMPLE1:
                 She was treated by her local doctor.
EXAMPLE1:
                 I’d like to make an appointment to see Dr Pugh.the doctor’s informal (=the place where your doctor works)
EXAMPLE1:
                 ‘Where’s Sandy today?’ ‘I think she’s at the doctor’s.’
Sense(doctor__2):2 someone who holds the highest level of degree given by a university → doctoral
EXAMPLE1:
                 a Doctor of Law
Sense(doctor__3):3  → be just what the doctor ordered
==find an ldoce entry==

HEAD doctor2 verb [transitive]
                 
                 
Sense(doctor__6):1 to dishonestly change something in order to gain an advantage
EXAMPLE1:
                 He had doctored his passport to pass her off as his daughter.
EXAMPLE1:
                 There are concerns that some players have been doctoring the ball.                                



Sense(doctor__7):2 to add something harmful to food or drink
EXAMPLE1:
                 Paul suspected that his drink had been doctored.
Sense(doctor__8):3 to remove part of the sex organs of an animal to prevent it from having babies SYN neuter
EXAMPLE1:
                 You should have your cat doctored.
Sense(doctor__9):4 to give someone medical treatment, especially when you are not a doctor
EXAMPLE1:
                 Bill doctored the horses with a strong-smelling ointment.
==find an bussdict entry==

HEAD doc‧tor1 /ˈdɒktəˈdɑːktər/ noun [countable]
Sense(doctor__11):1someone who is trained to treat sick or injured people
Sense(doctor__12):2someone who has a DOCTORATE from a university
==find an bussdict entry==

HEAD doctor2 verb [transitive]
Sense(doctor__14): to change something, especially in order to deceive people
EXAMPLE1:Police uncovered 43 cars for sale with doctored mileage readings.
]]
vimutil.open_floating_preview(vimutil.convert_input_to_markdown_lines(input), "markdown", {})
