package inet;

import java.io.IOException;
import java.net.InetAddress;
import java.net.UnknownHostException;

public class iptoname {
	/**
	 * @param args
	 */
	public static void main(String[] args) {
		//TODO Auto-generated method stub
		String IP =null;
		InetAddress host;//此类可以获取本机名和本机地址
		try{
			host = InetAddress.getLocalHost();//获取本机实例
			
			String localname = host.getHostName();//本机名
			String localip = host.getHostAddress();//本机地址
			
			System.out.println("本机名："+ localname  + "本机IP地址 " + localip);
			
		}catch(UnknownHostException e ){
			e.printStackTrace();
		}
		for(int i = 0;i<200;i++){
			IP = "192.168.2."+ i;
			try{
				host = InetAddress.getByName(IP);//获取IP的封装对象
				if(host.isReachable(2000)){//两秒来测试
					String hostName = host.getHostName();//获取制定IP地址的主机名
					System.out.println("IP地址"+IP+"的主机名称"+hostName);//
				}				
			}catch (UnknownHostException e) {//未知主机异常
				e.printStackTrace();
			}catch (IOException e) {//输入输出异常
				e.printStackTrace();
			}	
		}
		System.out.println("搜索完毕");
		}
}

