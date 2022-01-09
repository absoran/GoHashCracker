# GoHashCracker
# Introduction
This purpose of this project is crack the most commonly used hash algorithms. These algorithms are; SHA1,SHA256,SHA512 and MD5. I implement this 4 algorithm because these algorithms comes with crypto library in go base library. Application works on **CLI** (Commend Line Interface) and also with **Web**. 


# Used Technologies

- [Go](https://go.dev/) 
- JSON Web Token ([JWT](https://jwt.io/)) in authentication with this [library](https://github.com/dgrijalva/jwt-go)
- Used [PostgeSQL](https://www.postgresql.org/)'s cloud version [ElephantSQL](https://www.elephantsql.com/) with this [driver](https://github.com/lib/pq)
- Hypertext Markup Language (HTML)
- Cascading Style Sheets (CSS)

## Manual

The application works in 2 different modes. These modes are CLI mode and Web mode. In CLI mode user define and set flags. This flags are; mode for encryption or decryption. In decryption mode, user can must set --method flag and --hash flag. Application will use default wordlist. When user want to use different wordlist than default one, user must to set --haswordlist flag to true and properly set the --filepath flag with location of custom wordlist.

### CLI
- Encryption;
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929762669311963166/unknown.png" alt="logo" width="100%">

- Crack with default wordlist;
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929760004058259476/unknown.png" alt="logo" width="100%">

- Crack with custom wordlist;
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929761493795344415/unknown.png" alt="logo" width="100%">

- Crack hash with rule;
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929766852689801266/unknown.png" alt="logo" width="100%">
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929766513710354532/unknown.png" alt="logo" width="100%">

### Web
when user want to use web mode --mode flag must be set to web then web pages will be reachable

<img src="https://cdn.discordapp.com/attachments/235414073024446464/929768373041119302/unknown.png" alt="logo" width="50%">

- Signin page
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929770835043029002/unknown.png" alt="logo" width="80%">

- Signup page
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929770911341609040/unknown.png" alt="logo" width="80%">

- Homepage
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929771217358057512/unknown.png" alt="logo" width="80%">
<img src="https://cdn.discordapp.com/attachments/235414073024446464/929771361973469224/unknown.png" alt="logo" width="80%">
