package goserver

import (
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("User Repo powered by sqlite3", func() {
	var (
		db  *sqlx.DB
		err error
	)

	BeforeEach(func() {
		db, err = sqlx.Open("sqlite3", ":memory:")
		Expect(err).To(BeNil())
		db.Query(".read db_setup.sql")
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Creating a session repo", func() {
		It("Should have no error", func() {
			_, err := NewUserRepoSqlite3(db)
			Expect(err).To(BeNil())
		})
	})

	Context("Given an open repo", func() {
		var (
			u    *User
			repo *UserRepoSqlite3
			err  error
		)

		BeforeEach(func() {
			repo, err = NewUserRepoSqlite3(db)
			Expect(err).To(BeNil())

			u = &User{
				Username: "dummy",
				Password: "oh",
				Salt:     "pepper",
			}
		})

		Describe("Creating a user", func() {
			It("Should create a user", func() {
				var uu []*User

				err = repo.CreateUser(u)
				Expect(err).To(BeNil())

				err = db.Select(uu, "SELECT * FROM sessions")
				Expect(err).To(BeNil())
				Expect(len(uu)).To(Equal(1))
				u.ID = 1
				Expect(uu[0]).To(Equal(u))
			})
		})

		BeforeEach(func() {
			err = repo.CreateUser(u)
			Expect(err).To(BeNil())
		})

		Context("given an existing user", func() {
			var (
				uu []*User
			)

			Describe("Changing a user's password", func() {
				It("set the existing user's password", func() {

					err = repo.UpdateUserPasswd(u)
					Expect(err).To(BeNil())

					err = db.Select(uu, "SELECT * FROM seuuions")
					Expect(err).To(BeNil())
					Expect(len(uu)).To(Equal(1))
					u.ID = 1
					Expect(uu[0]).To(Equal(u))
				})
			})

			Describe("Authorize a user", func() {
				It("should get back the existing user", func() {
					err = repo.GetUser(u)
					Expect(err).To(BeNil())
					Expect(len(uu)).To(Equal(1))
					u.ID = 1
					Expect(uu[0]).To(Equal(u))
				})
			})
		})

	})
})
