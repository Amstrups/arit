f = open("table.html", "r")
table = f.read()

headerStart = table.find("<thead>")
headerEnd = table.find("</thead>")

if (headerEnd < 0):
    raise "fuck"

headers = table[headerStart:headerEnd]
shrinkStart = headers.find("<th>")
shrinkEnd = headers.rfind("</th>")
print(headers)
headers = headers[shrinkStart:shrinkEnd]
print(headers)


