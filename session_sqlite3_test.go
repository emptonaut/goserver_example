package goserver

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session Repo powered by sqlite3", func() {

	var (
		db  *sqlx.DB
		err error
	)

	BeforeEach(func() {
		db, err = sqlx.Open("sqlite3", "dummy.db")
		Expect(err).To(BeNil())
		//db.Query(".read db_setup.sql")
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Creating a session repo", func() {
		It("should return no error", func() {
			_, err := NewSessionRepoSqlite3(db)
			Expect(err).To(BeNil())
		})
	})

	Context("Given an open repo", func() {
		var (
			s    *Session
			repo *SessionRepoSqlite3
		)

		BeforeEach(func() {
			repo, err = NewSessionRepoSqlite3(db)
			Expect(err).To(BeNil())

			s = &Session{
				UserID:  41,
				Token:   "stupid",
				Origin:  "nowhere",
				Expires: "Always",
			}
		})

		Describe("Creating a session", func() {
			It("should create a good session", func() {
				var ss []*Session

				err = repo.CreateSession(s)
				Expect(err).To(BeNil())

				err = db.Select(ss, "SELECT * FROM sessions")
				Expect(err).To(BeNil())
				Expect(len(ss)).To(Equal(1))
				s.ID = 1
				Expect(ss[0]).To(Equal(s))
			})
		})

		Context("given an existing session", func() {
			var (
				ss []*Session
			)

			BeforeEach(func() {
				err = repo.CreateSession(s)
				Expect(err).To(BeNil())
			})

			Describe("Deleting a session", func() {
				It("Should delete", func() {
					err = repo.DeleteSession(s)
					Expect(err).To(BeNil())
					err = db.Select(ss, "SELECT * FROM sessions")
					Expect(err).To(BeNil())
					Expect(len(ss)).To(Equal(0))
				})
			})

			Describe("Retrieve a session", func() {
				It("should get back the existing session", func() {
					err = repo.GetSession(s)
					Expect(err).To(BeNil())
					Expect(len(ss)).To(Equal(1))
					s.ID = 1
					Expect(ss[0]).To(Equal(s))
				})
			})
		})
	})
})
