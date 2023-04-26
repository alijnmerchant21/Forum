package forum

/* Explanation

* User struct: Has three fields: name, banned, and bannedPost. The name field stores the username of the user,
the banned field stores a boolean value indicating whether the user is banned, and the bannedPost field is a map that stores the IDs of posts
that the user has had banned.

* NewUser function: Is a constructor that takes a username as a parameter and returns a new User object with the name field set to the
given username and the banned field set to false. The bannedPost map is initialized to an empty map using the make function.

* AddPost method: Takes a post ID as a parameter and adds it to the bannedPost map, with the value set to false.
If the user is currently banned, the post is not added.

* Ban method: Sets the banned field to true, indicating that the user has been banned.

* IsBanned method: Returns a boolean value indicating whether the user is currently banned.

* BanPost method: Takes a post ID as a parameter and sets the value in the bannedPost map to true, indicating that the post has been banned.
If the user is currently banned, the post is not banned.

* IsPostBanned method: Takes a post ID as a parameter and returns a boolean value indicating whether the post has been banned for the user.

* BanUser function: Sets the banned field to true, indicating that the user has been banned.

* AddUser function: Takes a username as a parameter and updates the name field of the User object to the given username.
It also sets the banned field to false and initializes the bannedPost map to an empty map using the make function.
*/

type User struct {
	name       string
	banned     bool
	bannedPost map[int]bool
}

func NewUser(name string) *User {
	return &User{
		name:       name,
		banned:     false,
		bannedPost: make(map[int]bool),
	}
}

func (u *User) AddPost(postID int) {
	if u.banned {
		return
	}
	u.bannedPost[postID] = false
}

func (u *User) Ban() {
	u.banned = true
}

func (u *User) IsBanned() bool {
	return u.banned
}

func (u *User) BanPost(postID int) {
	if u.banned {
		return
	}
	u.bannedPost[postID] = true
}

func (u *User) IsPostBanned(postID int) bool {
	return u.bannedPost[postID]
}

func (u *User) BanUser() {
	u.banned = true
}

func (u *User) AddUser(name string) {
	u.name = name
	u.banned = false
	u.bannedPost = make(map[int]bool)
}
