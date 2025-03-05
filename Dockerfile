FROM golang:1.23.4

LABEL Name="forum" \
    Version="Latest" \
    Contributors="Allan Kamau <https://learn.zone01kisumu.ke/git/allkamau>, Fred Gitonga <https://learn.zone01kisumu.ke/git/fgitonga>, Ann Maina <https://learn.zone01kisumu.ke/git/nymaina>, cheryl owalla <https://learn.zone01kisumu.ke/git/cowalla>"

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o /forum

EXPOSE 9000

CMD ["/forum"]
