### Tela (Home)

- [x] user/info/previewUsers

```bash
# Parâmetros
$ userId
```

- [x] search/content/results

```bash
# Parâmetros
$ querySearch
$ pageDest
$ forceResults
```

### Tela (Cliente)

- [x] user/info/dectecUser

```bash
# Parâmetros
$ msisdn
```

- [x] user/info/discountFidelization

```bash
# Parâmetros
$ msisdn
```

- [x] user/fidelization/info

```bash
# Parâmetros
$ msisdn
```

- [x] user/info/contractDetails

```bash
# Parâmetros
$ msisdn
```

- [x] user/plan/consult

```bash
# Parâmetros
$ channel
$ commercialCode
```

- [x] user/protocol/fullHistory

```bash
# Parâmetros
$ msisdn
$ startDate
$ endDate
$ number
```

### Componente Card de Informações Importantes

- [x] user/ura/lastcal

```bash
# Parâmetros
$ msisdn
```

- [x] user/balance/info

```bash
# Parâmetros
$ msisdn
```

- [x] user/invoice/info

```bash
# Parâmetros
$ msisdn
```

- [x] user/info/eligibilityC6

```bash
# Parâmetros
$ socSecNo
```

- [x] search/procedure/messages

```bash
# Parâmetros
$ assetId
$ planName
$ segmentId
$ state
```

- [x] conversations

```bash
# Parâmetros
$ dinâmicos
```

- [x] user/invoice/billingRuler

```bash
# Parâmetros
$ customerId
```

- [x] user/invoice/billingRulerHistory

```bash
# Parâmetros
$ customerId
$ segmentId
$ initialDate
```

- [x] user/invoice/firstInvoice

````bash
# Parâmetros
$ contractId
$ msisdn

- [x] user/protocol/portabilityHistory

```bash
# Parâmetros
$ msisdn
````

- [x] CMS/priceUp

```bash
# Parâmetros
$ msisdn
```

### Card consumo de voz / Dados

- [x] user/consumption/data

```bash
# Parâmetros
$ msisdn
$ segmentId
```

- [x] user/consumption/voice

```bash
# Parâmetros
$ msisdn
```

### Card Histórico de Atendimento

- [x] user/protocol/fullHistory

```bash
# Parâmetros
$ msisdn
$ startDate
$ endDate
$ number
```

### Card Histórico de Faturas

- [x] user/invoice/infoHistory

```bash
# Parâmetros
$ msisdn
```

### Card de Histórico de Recarga

- [x] user/balance/infoHistory

```bash
# Parâmetros
$ msisdn
```

### Card Últimas Movimentações

- [x] user/balance/infodetail

```bash
# Parâmetros
$ msisdn
$ dateStart
$ dateEnd
$ mms
$ pageNumber
$ pageSize
$ recharge
$ refund
$ sms
$ voice
$ data
$ vas
$ events
```

### Seção Contratos

- [x] user/promotion/info

```bash
# Parâmetros
$ msisdn
$ lineType
```

- [x] user/info/contractSummary

```bash
# Parâmetros
$ msisdn
$ userId
```

### Seção Visita Técnica

- [x] user/live/visitStatus

```bash
# Parâmetros
$ socialSecNo
$ installationNumber
$ status
```

### Seção Cliente

- [x] user/relation/info

```bash
# Parâmetros
$ msisdn
```

- [x] search/notifications/messages

```bash
# Parâmetros
$ Sem parâmetros
```

- [x] user/protocol/recentContacts

```bash
# Parâmetros
$ msisdn
$ cpf
$ segmentId
```

### Menus no canto superior direito

#### Essas API's só são requisitadas quando clicado no item respectivo do menu

- [x] CMS/procedimento
- [x] CMS/message
- [x] CMS/segmentos
- [x] search/systemsMenus
- [x] search/featuredProcedures
