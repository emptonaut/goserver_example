package goserver

import (
	"fmt"

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

		AfterEach(func() {
			_, err = db.Exec("DELETE FROM sessions")
			Expect(err).To(BeNil())
			_, err = db.Exec("delete from sqlite_sequence where name='sessions';")
			Expect(err).To(BeNil())
		})

		Describe("Creating a session", func() {
			It("should create a good session", func() {
				var ss []*Session

				err = repo.Create(s)
				Expect(err).To(BeNil())

				err = db.Select(&ss, "SELECT * FROM sessions")
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

			JustBeforeEach(func() {
				err = repo.Create(s)
				Expect(err).To(BeNil())
				Expect(s.ID).To(Equal(1))
			})

			Describe("Deleting a session", func() {
				It("Should delete", func() {
					err = repo.Delete(s)
					Expect(err).To(BeNil())
					err = db.Select(&ss, "SELECT * FROM sessions")
					Expect(err).To(BeNil())
					Expect(len(ss)).To(Equal(0))
				})
			})

			Describe("Retrieve a session", func() {
				It("should get back the existing session", func() {
					fmt.Println(s)
					uut := &Session{ID: 1}
					err = repo.GetByID(uut)
					Expect(err).To(BeNil())
					s.ID = 1
					Expect(uut).To(Equal(s))
				})
			})
		})
	})
})
