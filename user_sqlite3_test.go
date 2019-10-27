package goserver

import (
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Repo powered by sqlite3", func() {
	var (
		db  *sqlx.DB
		err error
	)

	BeforeEach(func() {
		db, err = sqlx.Open("sqlite3", "dummy.db")
		Expect(err).To(BeNil())
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

		AfterEach(func() {
			_, err = db.Exec("DELETE FROM users ")
			Expect(err).To(BeNil())
			_, err = db.Exec("delete from sqlite_sequence where name='users';")
			Expect(err).To(BeNil())
		})

		Describe("Creating a user", func() {
			It("Should create a user", func() {
				var uu []*User

				err = repo.CreateUser(u)
				Expect(err).To(BeNil())

				err = db.Select(&uu, "SELECT * FROM users")
				Expect(err).To(BeNil())
				Expect(len(uu)).To(Equal(1))
				u.ID = 1
				Expect(uu[0]).To(Equal(u))
			})
		})

		Context("given an existing user", func() {
			var (
				uu []*User
			)

			BeforeEach(func() {
				err = repo.CreateUser(u)
				Expect(err).To(BeNil())
				Expect(u.ID).To(Equal(1))
			})

			Describe("Changing a user's password", func() {
				It("set the existing user's password", func() {

					err = repo.UpdateUserPasswd(u)
					Expect(err).To(BeNil())

					err = db.Select(&uu, "SELECT * FROM users")
					Expect(err).To(BeNil())
					Expect(len(uu)).To(Equal(1))
					u.ID = 1
					Expect(uu[0]).To(Equal(u))
				})
			})

			Describe("Retrieve a user", func() {
				It("should get back the existing user by ID", func() {
					u.ID = 1
					uut := &User{ID: 1}
					err = repo.GetUserByID(uut)
					Expect(err).To(BeNil())
					Expect(uut).To(Equal(u))
				})
				It("should get back the existing user by username", func() {
					u.ID = 1
					uut := &User{Username: "dummy"}
					err = repo.GetUserByUsername(uut)
					Expect(err).To(BeNil())
					Expect(uut).To(Equal(u))
				})
			})
		})
	})
})
