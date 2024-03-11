# FSA - Fundamentals of Software Architecture 
##### An Engineering Approach

By Mark Richards and Neal Ford

## Notes:

## Self-Assessment Questions:

### Chapter 1: Introduction:
1. What are the four dimensions that define software architecture?
    - Architecture consists of the ***STRUCTURE*** combined with ***ARCHITECTURE CHARACTERISTICS***, ***ARCHITECTURE DECISIONS*** and ***DESIGN PRINCIPLES***.
2. What is the difference between an architecture decision and a design principle?
    - Architecture decision is strict rule defined how the system should be constructed. For example, we have layered architecture style, we decided to follow the closed style, any layer shouldn't be able to pass through layers without the call the next beneath. Follow this rule is not recommendation! Is rule and should be followed. On the other hand, a design principle is used to provide guidance or a recommendation from the architecture to the developer team. For example, "whenever possible, leverage async messaging between services to increase performance"
3. List the eight core expectations of a software architect.
    - Make architecture decisions, continually analyze the architecture, keep current with the latest trends, ensure compliance with decisions, diverse exposure and experience, have business domain knowledge, possess interpersonal skills and understand and navigate politics
4. What is the first Law of Software Architecture?
    - There are no right or wrong answers in architecture only trade-offs.
### Chapter 2: Architecture Thinking
1. Describe the traditional approach of architecture versus development and explain why that approach no longer works.
    - The traditional approach is very similar with the waterfall project management, where the architecture do his work about the requirements and once done he passes to the development team what should be done and don't iterate about the architecture with the development team again. This causes a unidirectional arrow of interactions, resulting in architecture decisions that will never be developed by the team and some development decisions will never change the architecture.
2. List the three levels of knowledge in the knowledge triangle and provide and examples of each.
    - Stuff you, example: is the technologies/knowledge that you use in your daily basis to do your job. Stuff you know you don't know, example: is the technologies/knowledge that you have heard about it but never used or just know that exists. Stuff you don't know you don't know, example: is all the technologies/knowledge that could solve the problem that you are facing, but you even know that such thing exists.
3. Why is it more important for an architect to focus on technical breadth rather than technical depth?
    - Because an architect must make decisions that match capabilities to technical constraints. For example is much more important to an architect to know that exist 7 types of cache products and their uses cases and pros and cons, that be a specialist at one cache product, each project has different requirements and problems to solve, so he must know each product fits better in these constraints.
4. What are some of the ways of maintaining your technical depth and remaining hands-on as an architect?
    - Proof-of-concepts(POCs), tackling the low priority user stories or work on bug fixes both in an iteration and do code reviews.
### Chapter 3: Modularity
1. What is meant by the term connascence?
    - "Two components are connascent if a change in one would require the other to be modified in order to maintain the overall correctness of the system" - Meilier Page-Jones(1996-Dorset-House)
2. What is the difference between static and dynamic connascence?
    - Static connascence refers to the source-code level coupling, for example multiple components depends on a specific entity type. And dynamic connascence refers to the connascence in the execution time, for example the order execution of a method is a dynamic connascence.
3. What does connascence of type mean? Is it static or dynamic connascence?
    - Is a static connascence. Connascence of type is when multiple components depends on a specific entity type in order to work
4. What is the strongest form of connascence?
    - Dynamic connascence of type identity, when multiple system must reference to the same entity in the execution time.
5. What is the weakest form of connascence?
    - Static connascence of type name, when multiple components must agree on the name of an entity.  
6. Which is preferred within a code base-static or dynamic connascence?
    - Static connascence, because the static connascence is much more in the control of the architect and developer team to refactor to a better type of coupling than dynamic, nowadays, the lack of tools to measure dynamic connascence is high and analyze e measure dynamic connascence is much harder than static.
### Chapter 4: Architecture Characteristics Defined
1. What three criteria must an attribute meet to be considered an architecture characteristic?
    - Specifies a nondomain design consideration; Influences some structural aspect of the design; Is critical or important to application success.
2. What is the difference between an implicit characteristic and an explicit one? Provide an example of each
    - Implicit architecture characteristics is the ones that rarely appear in the requirements document, yet they are necessary for project success. For examples, security no one wants a vulnerable application but no one specifies how secure should be the app.
3. Provide an example of an operational characteristic.
    - Operational characteristic is all characteristics that a system will have to ensure responsiveness to the client an example: Recoverability
4. Provide an example of a structural characteristic.
    - Structural characteristic is all characteristics that ensure code quality an example: Maintainability and Extensibility
5. Provide an example of a cross-cutting characteristic.
    - Cross-cutting characteristic is all characteristics that ensure security an example: Privacy and Authorization.
6. Which architecture characteristic is more important to strive for availability or performance?
    - It depends on the requirements, if for the business is more important the system is need to be available 24/7 in case in any failure, so we should strive for availability, but if it is capacity of low response time in any case should be performance.
### Chapter 5: Identifying Architecture Characteristics
1. Give a reason why it is good practice to limit the number of characteristics("-ilities") an architecture should support.
    - If an architecture does all characteristics she does nothing very well, this will lead in a generic solution means that architecture will not success in the core requirements for the business. This repents because each architecture characteristics impact on each other for example when you choose to support security probably you will have to trade of by performance because you probably you're going to have to do more encryption for each data response because encryption is expensive and takes time.
2. True or False: most architecture characteristics come from business requirements and user stories.
    - False, most architecture characteristics comes from the architect interpretation of the requirements document or business stakeholder conversation. 
3. If a business stakeholder state that time-to-market(i.e., getting new feature and bug fixes pushed out to users as fast as possible) is the most important business concern which architecture characteristics would the architecture need to support?
    - Testability+deployability+agility.
4. What is the difference between scalability and elasticity?
    - Scalability is the ability to handle a large number of concurrent users without serious performance degradation, therefore elasticity is the ability to handle bursts of requests. Elasticity look for metrics like burst of users on time and scalability looks for performance of concurrent users. Some time a system can be Scalable but not elastic to handle bursts requests.
5. You find out that your company is about to undergo several major acquisitions to significantly increase its customer base. Which architectural characteristics should you be worried about?
    - Scalability the ability to handle many concurrent users on time. Also, I would be worried with Availability and deployability as architect I want to ensure that my system will not stop delivery feature because the amount of new users and the development process should not change because of that.
### Chapter 6: Measuring and Governing Architecture Characteristics
1. 