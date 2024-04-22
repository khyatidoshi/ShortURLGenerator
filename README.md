# URL-Shortener
This project is a URL shortener that accepts a URL of any length and generates a fixed-length short URL for it. When a short URL is accessed, the user is redirected to the long URL. Every time the Short URL is accessed, the event is stored in the Database until the Short URL expires. 

## Requirements
#### Functional
- The service should generate a unique short URL for the provided long URL
- The service should redirect the user to the original URL when the short link is called
- The short link should have a lifetime that the user specifies on creation or never expire
- The short link should track the click count

#### Non-functional
- The service should be able to handle numerous requests (Scalable)
- Forwarding should be real-time with minimum delay (Low latency)
- The short link should be random in order not to be predictable (Secure solution) 

## Estimations
The service is going to serve heavy reads since there will be a huge number of redirects compared to creating new ones, a good thing would be to have a cache. Let’s assume that the ratio between reading and writing is 50:1. 
##### Traffic estimates
If we have 500k new short links every month, then we will expect 25 million (50 * 500k = 25 million) redirects for the same period. So we have 1 new link every 5 seconds: 
`500k / (30 days * 24 hours * 3600 seconds) = ~ 1 link in 5 seconds.`

And 10 redirects every second: 
`25 million / (30 days * 24 hours * 3600 seconds) = ~ 10 redirects per second.`

##### Memory estimates
Let’s say we store each address for a maximum - 1 year. We expect 500k new links every month, and then we will have nearly 6 million records in the database: 
`500k record/month * 12 months = 6 million`

Let’s assume that each record in the database - approximately 1000 bytes. [The recommended maximum size for a link is 2000 characters](https://stackoverflow.com/questions/417142/what-is-the-maximum-length-of-a-url-in-different-browsers/417184#417184) and according to the standard, the URL encodes with ASCII characters, which occupy 1 byte, i.e. the link can hold  2000 bytes by recommended maximum size. So we will use half of this value as an average. Then we need 6 GB of memory to store records for 1 year: 
`6 million record * 1000 bytes per record = 6 GB`

>A little summary of the nature of the model:
>- We need to store several million records
>- Each record is small 
>- The service is very read-heavy

## Assumptions
* The expiration date for the URL can only be given at short URL generation and can't be changed later.
* Once the short URL is deleted, all events of when the URL was accessed will no longer be required.
* every second the application would receive 100 requests at the peak time.

## Features
* Url Shortener
* Tracking capabilities
  * Clicks on the link can be tracked.
* Delete short URL if the date of expiry is given at short URL creation
  
## 📚 | Introduction
- It uses Redis as a cache, and PostgresDB as an SQL database.
- It uses Nginx which acts as a load balancer and a reverse proxy for the backend server.
- It uses [Go's UUID package](https://pkg.go.dev/github.com/google/uuid) to generate an unique short URL.
## Design 
![Test Image 1](https://github.com/khyatidoshi/ShortURLGenerator/blob/main/ShortURLWorkflow.JPG)
<br>
<b> Unique Short URL Generation </b>: Go's UUID package func NewRandom Randomly generated UUIDs having 122 random bits. We slice this and take the first 8 characters for generating short links. total unique links that can be produced would be 36^8 = 2,821,109,907,456. If we generate 500 new short URLs per sec, we would consume 500X60 = 30,000 new links every minute, and 500X60X60 = 1,800,000 new links every hour. Consuming 36^8 links would take more than 100 years. The chances of generating a duplicate short link would be 1/36^8 which is extremely low. 
<br>
<b> Choice of Database </b>: The data in this project has a proper structure - we already know exactly what details to save, and have an exact structure for the data. MySQL is widely used and queries in Postgres are similar. We also need to consider the cost of managing the Database with Cloud services like AWS. 

## 🚀 | Usage

- Install Docker Desktop and run it.
- Clone this repository:<br>

```sh
git clone https://github.com/khyatidoshi/ShortURLGenerator.git
```

- Open the project folder and start the container with docker-compose:<br>

```yml
docker compose up --build

```
### Testing 
- Download Artillery:<br>
```sh
For global install --- npm install -g artillery
```
OR 
```sh
For local install --- npm install artillery
```

- Running the test:<br>
```sh
if globally installed ---  artillery run *filename*
```
OR 
```sh
if locally installed --- npx artillery run *filename*
```
