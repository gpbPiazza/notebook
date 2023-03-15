# DDD - Domain Driven Design

###  Putting the Domain Model to Work

-  A model is a simplification of reality. It is an interpretation of reality that abstracts the aspects relevant to solving the problem at hand and ignores extraneous detail

### One. Crunching Knowledge

- Build a diagram, with any annotation or any just have something em common with the expertise of domain to you communicate what have you understand about the domain so he can give you feedback about it.

- When we have an unimportant model and an important have an interaction we should build a new model different from the other ones Eric called this as *Distilling the model*, given a example :
    - Model A is relevant and model B not. We have a new use case where we should use model B with model A. So this new use case is a new model, therefore we should make a modelAB or modelC to make clear the relevance of model A stay relevant e not "contaminated" with the irrelevance of B.

- Build a model -> test -> feedback -> repeat (massive experimentation)

### Two. Communication and the Use of Language

- Ubiquitous Language! A project where we don't have the Ubiquitous Language present the developers have to translate what the specialist of domain are saying to others developers, the developers have to translate what they have into domain today to the specialists, all this translation only make miss understanding of what the domain should be, all this translation makes the team run slower and build a poor domain.

- Ubiquitous Language forces the weaknesses of domain out, makes the team communicate in one language and ensure with the language changes the code will too.

- The preparation to new models or refine the domain you must Modeling Out Loud! Explore with speech. Trying out loud various constructs from possible model variations. It's easy to hear what is wrong.

![obiquitous langue diagram](./assets/obiquitous-langue-diagram.png "obiquitous-langue-diagram")

- Documents and diagrams should be used to explain the main frame of domain and none functional requirements of a system. They should short, clear and concise.

- The fundamental details about the domain are in the code!

- Document everything? Or Anything? A document should not do what a code do well. The code give us all details about the system. Therefore, use diagrams to get the big picture of the system.