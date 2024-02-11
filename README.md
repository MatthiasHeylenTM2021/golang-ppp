# GO PPP - Go Playground Projects

Welcome to the GO PPP - Go Playground Projects! 🚀

This repository is a dynamic collection of hands-on projects and experiments, exploring various fundamental concepts of Golang. Dive in and explore the world of Go through the following folders:

## 1. hello-web-project
   - **Description:** This directory hosts the Hello Web Project, a dynamic exploration of Golang's capabilities in web development. From handling HTTP requests to diving into the essentials of web applications, this project serves as a starting point for anyone eager to get hands-on with Golang on the web.

## 2. bookings
   - **Description:** Behold, the Bookings directory! This is where the magic happens. Explore the intricate details of a Bed & Breakfast website built entirely with Golang. From handling reservations to managing the backend intricacies, this project encapsulates real-world application development in Golang.

## 3. concepts
   - **Description:** The Concepts directory is a treasure trove of small projects aimed at solidifying your understanding of Golang basics. From loops and ranging to decision structures, this collection provides bite-sized coding exercises to reinforce your knowledge and enhance your coding skills.

## Getting Started

Before diving into the projects, make sure you have the following tools and dependencies installed on your machine:

1. **Golang:** Ensure you have Golang installed on your machine. You can download it from [Golang Official Website](https://golang.org/dl/).

2. **PostgreSQL:** The projects might involve database interactions. Install PostgreSQL and set up a database. You can download it from [PostgreSQL Official Website](https://www.postgresql.org/download/).

   - **Example Connection String:**
     ```bash
     postgres://username:password@localhost:5432/database_name?sslmode=disable
     ```
     Make sure to replace `username`, `password`, and `database_name` with your PostgreSQL credentials.

3. **MailHog:** To simulate email interactions, use MailHog. Download and install it from [MailHog GitHub Releases](https://github.com/mailhog/MailHog/releases).

4. **DBeaver:** For database management and visualization, DBeaver is a handy tool. Download it from [DBeaver Official Website](https://dbeaver.io/download/).

5. **Soda/Pop from Buffalo:** If you are dealing with migrations, Soda (or Pop) is a powerful tool from the Buffalo framework. Find more information on [Buffalo Pop Documentation](https://pkg.go.dev/github.com/gobuffalo/pop).

   - **Installation:**
     ```bash
     go get -u github.com/gobuffalo/pop/...
     ```

6. **Foundation for Email Templates:** For email templates, you can use the Foundation framework. Check out the available templates at [Foundation Email Templates](https://get.foundation/emails/email-templates.html).

7. **Simple DataTables for Admin Dashboard:** For a clean table in the Admin Dashboard, consider using Simple DataTables. Get started with DataTables at [DataTables Official Website](https://datatables.net/).

   - **Example Integration:**
     ```html
     <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.11.5/css/jquery.dataTables.css">
     <script type="text/javascript" charset="utf8" src="https://code.jquery.com/jquery-3.6.0.js"></script>
     <script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.11.5/js/jquery.dataTables.js"></script>
     ```

8. **Soda Migrate Commands:**
   - **Run Migrations:**
     ```bash
     soda migrate up
     ```
   - **Generate a New Migration:**
     ```bash
     soda generate fizz MyNewMigration
     ```
   - **Rollback Last Migration:**
     ```bash
     soda migrate down
     ```

Now you're all set! Clone this repository to your local environment, navigate to the specific project directory you want to explore, and start tinkering, experimenting, and learning!

Feel free to customize the database connection string, email templates, DataTables integration, and Soda migrate commands based on your preferences and project requirements.

Happy coding! 🚀
