FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD login login
ADD static static
ADD views views
EXPOSE 8080
CMD ["/login"]
