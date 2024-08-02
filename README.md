scanNetwork fonksiyonu, nmap komutunu çalıştırarak ağdaki cihazların IP adreslerini döner.
sendMessage fonksiyonu, belirli bir IP adresine HTTP POST isteği gönderir.
main fonksiyonu, ağı tarar ve bulunan tüm cihazlara mesaj gönderir.
Not: Bu örnekte, telefonların 8080 portunda dinlediği varsayılmaktadır. Gerçek uygulamada, telefonların hangi portta dinlediğini ve nasıl mesaj kabul ettiğini bilmeniz gerekir. Ayrıca, güvenlik ve yetkilendirme konularını da göz önünde bulundurmalısınız.
