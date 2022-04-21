## Automate Google Calendar Bulk Invites : Powered by Golang
This tool levarages Google Calendar API and automate calendar events creation to specified guests with a specified event start and end time.

Use Case
- Whenever you need to send out Calendar Invites to a group of users
- Why not to use Google Calendar API instead of GUI

## Usage

- Get the code
```
git clone https://github.com/ksingh7/google-calendar-automated-events-go
cd google-calendar-automated-events-go
```
- Create Google Cloud Platform project with the Google Calendar API enabled [follow this guide](https://developers.google.com/workspace/guides/create-project)
  - go to `API & Services` -> `Library` -> `Google Calendar API` -> `Enable`
- Authorization credentials for a desktop application (oauth-client-id) [follow this guide](https://developers.google.com/workspace/guides/create-credentials#oauth-client-id)
  - go to `API & Services` -> `Credentials` -> `Create Credentials` --> `oauth-client-id` -> `Desktop App` -> `Download Json` -> save as `credentials.json` in the code directory
- From project directory run `go mod tidy` to update the dependencies
- Update `talk_details.csv` with details of calendar invites
- Execute `go run main.go`
- Sample Output
```
$ go run main.go
&{Getting a head start in career with Kubernetes 2022-04-23T11:00:00+05:30 2022-04-23T11:15:00+05:30 cnu1812@gmail.com}
Email,  karan.singh731987@gmail.com
Talk Title,  Getting a head start in career with Kubernetes
Talk Start Time,  2022-04-23T11:00:00+05:30
Talk End Time,  2022-04-23T11:15:00+05:30
Event created: https://www.google.com/calendar/event?eid=M2pmMGpiZjdhaXBjYmFoOWtuZDc3YWlyY2Mga2FyYXNpbmdAcmVkaGF0LmNvbQ
&{Running local Kubernetes clusters using minikube, kind and microk8s 2022-04-23T11:15:00+05:30 2022-04-23T11:30:00+05:30 mehabhalodiya@gmail.com}
Email,  karan.singh731987@gmail.com
Talk Title,  Running local Kubernetes clusters using minikube, kind and microk8s
Talk Start Time,  2022-04-23T11:15:00+05:30
Talk End Time,  2022-04-23T11:30:00+05:30
Event created: https://www.google.com/calendar/event?eid=NHJiNXE0YmM5c2tnbjdhZ3Q0MDYzaDRkMGcga2FyYXNpbmdAcmVkaGF0LmNvbQ
&{Orchestrating Cloud Native ML workflows in Kubernetes 2022-04-23T11:30:00+05:30 2022-04-23T12:00:00+05:30 senthilrch@gmail.com}
$
```