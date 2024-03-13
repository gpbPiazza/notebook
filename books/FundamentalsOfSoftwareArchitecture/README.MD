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
1. Why is cyclomatic complexity such an important metric to analyze for architecture?
    -  Cyclomatic complexity is metric to analyze code structure complexity where is verified cyclomatic complexity of a code using cc = E-N+2P where E represents edges(if statements count) and N represents nodes (lines of code) and P represents the number of reference for that method or function. Using this metric the architects and developers can common agree if a class or function is too complexity or not, too much complexity can mean that some architecture characteristics is not been satisfied such as testability or modularity.
2. What is an architecture fitness function? How can they be used to analyze an architecture?
    - Any mechanism that provides an objective integrity assessment of some architecture characteristic or combination of architecture characteristics. They can be used do ensure compliance of an architecture characteristic for example we could write a unit tests verifying the in dependency between architecture layers and if they are satisfying the architecture design expected. Others examples, metrics, chaos engineering and monitors.
3. Provide an example of an architecture fitness function to measure the scalability of an architecture.
    - You should define performance of your application and a range of it, for examples performance for me is the response time of each endpoint exposure in the API and I will get the simple median from this and say it is the overall performance of my application for examples high performance 300 ms and low performance 1000 ms. Ok now I will join a metric of concurrent users, such as quantity of request on small range of time and will define a monitor to verify my median performance by the concurrent users metric, if this metric reaches some kind of value of low performance if high concurrent users I will make this monitor notify my team to take a look in the system.
4. What is the most important criteria for an architecture characteristic to allow architects and developers to create fitness functions?
    - Choose the architecture characteristics that your application will satisfy the most and by fitness functions ensure the compliance.
### Chapter 7: Scope of Architecture Characteristics
1. What is an architecture quantum, and why is it important to architecture?
    -  Architecture quantum is - "An independently deployable artifact with high functional cohesion and synchronous connascence", where "independently deployable artifact" means all structure necessary to say that a system is health for examples (a monolith and his DB), and "high functional cohesion" means how well contained code is unified in purpose, and "cohesion and synchronous connascence" means sync calls within an application context or between distributed services(synchronous connascence = coupling between services in execution time by requests).
    - Architecture quantum is important because give the architects a new metric to analyze the high level architecture design, many architecture quantum indicates a distributed system, and we can assume a lot just know this kind of number. For example if a system only has 1 architecture quantum certainly this system is a monolith application
2. Assume a system consisting of single user interface with four independently deployed services, each containing its own separate database. Would this system have a single quantum or four quanta? Why?
    - Four quanta, because each separated deployable unit artifact, the independently deployed services, satisfies one quantum each. Someone also could say that this architecture is five quantum, counting the frontend as one quantum, but the front end is not an independently deployed system he depends on four services to work on, so he doesn't count in the architecture quantum count.
3. Assume a system with an administration portion managing static reference data (such as product catalog and warehouse information) and customer facing portion managing the placement of orders. How many quanta should this system be and why? If you envision multiple quanta, could the admin quantum and customer facing quantum share a database? If so, in which quantum would the database need to reside?
    - If we assume the admin system and the customer system are independently deployable units therefore this architecture has two quanta. Each has distinct responsibilities and could evolve independently. While the two quanta could technically share a database, it would be better for each to have its own to maintain the principle of independent deployability. Sharing a database would introduce a level of coupling that could inhibit the ability to deploy and scale each quantum independently. If they were to share a database, it would reside in neither quantum specifically, but would become a shared resource, which would need to be carefully managed to avoid conflicts and ensure consistency.
### Chapter 8: Component-Based Thinking
1. We define the term component as a building block of an application something the application does. A component usually consist of a group of classes or source files. How are components typically manifested within an application or service?
    - Components are usually manifested within an application or service as distinct units of functionality. They can be organized as modules, packages, or services, each encapsulating a set of related functionalities. In object-oriented languages, a component can comprise multiple classes, interfaces, or packages. In a microservices architecture, each microservice can be considered a component.
2. What is the difference between technical portioning and domain partitioning? Provide an example of each.
    - Technical portioning is a way to organize the source code by technical capabilities, for example a layered architecture presentation layer, business layer, persistence layer and so on. It makes clear that all code to lead with the database is inserted in the persistence layer. The domain portioning is a way to organize the source code by workflows, or business domain functionalities, where each part is bounded context of the system domain, for examples microservices is a architecture style where each service has his own domain. The main difference between this portioning is the top level organization.
3. What is the advantage of domain partitioning?
    - Partitioning a system by domains better reflects the kinds of changes that system will need when business rules change, models the system more close how the business functions work and more easy to migrate to a distributed architecture.
4. Under what circumstances would technical partitioning be a better choice over domain partitioning?
    - Technical partitioning is often a patter used to starter projects or business that don't have all knowledge of his domain, offering a low learning curve for developer to iterate in the code base and learn about the business. Also in an organization where exists specialists teams like the front end team, the backend team and DB teams this kind of organization better fits into technical partitioning.
5. What is the entity trap? Why is it not a good approach for component identification?
    - Entity trap is antipattern when an architect incorrectly identifies the database relationships as workflows in the application, this correspondence that rarely manifests in the real world. Components created by this pattern will result in coarse-grained components where it will not give any useful guidance to developers team.
6. When might you choose the workflow approach over the Actor/Actions approach when identifying core components?
    - 
### Chapter 9: 