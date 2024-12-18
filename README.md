## Posadas Reports
Posadas Reports is an application designed to help residents of Posadas (Argentina) report various issues related to public safety, environmental concerns, and urban maintenance. The app provides a map-based interface for users to report problems such as unsafe neighborhoods, poor lighting, environmental hazards, issues with waste management and recycling, etc.

### Features
- Geographical Reporting: Users can select a point on a map to report issues specific to that area.
- Categories: Reports can be categorized by type, such as environmental problems, safety concerns, and waste disposal issues and by urgency, being Low, Medium and Critical some of the options available.
- No Registration: No need to sign up or connect to your Google account.
- Statistics: Provides insights into the most frequent report categories and the number of recent reports to identify recurring issues.

### Technology Stack
- Go 
- Gin
- GORM
- PostgreSQL (formerly SQLite)
- TEMPL 

### Usage
To get started with the app:

- Navigate to the map interface.
- Select the region where the issue has occurred.
- Choose the appropriate category for the issue.
- Enter a title and description of the problem.
- Define the report's urgency (urgency criteria below this section).
- Submit the report.

### Urgency criteria
- None: The issue does not pose any immediate risk. If left unaddressed, there will be no significant consequences or impact on public safety or the environment.
- Low: The problem is likely to cause minor inconvenience or annoyance to passers-by, but it does not pose any significant threat. While it may lead to discomfort or aesthetic concerns, there is no immediate danger to people or property.
- Medium: There is a potential for minor accidents or disruptions. The issue could lead to situations that may cause injury or damage if left unchecked, though the risk remains relatively low. Immediate attention is recommended to prevent escalation.
- High: The issue presents a significant risk to public safety or property. There is a real potential for serious injury to individuals or damage to both private and public property. Delaying action could result in more severe consequences or widespread impact.
- Critical: The situation is currently causing harm to individuals or extensive damage to property. Immediate intervention is necessary to prevent further injury, loss of life, or irreversible damage. This is an emergency-level concern requiring urgent resolution to safeguard the well-being of the community.

### License
This project is licensed under the MIT License - see the LICENSE file for details.