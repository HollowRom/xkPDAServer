use AIS20220623135859


select top 10 *
from xkPdaServer_good_tool
where FMATERIALID = 118405 --and FUSEORGID = 1 and FISINVENTORY = 1
select top 10 *
from t_BD_MaterialStock
where FMATERIALID = 118405
select top 10 *
from T_BD_UNIT
where FMASTERID = 100007
select top 10 a.FMASTERID            as FMATERIALID,    --物料内码
              a.FNUMBER,                                --物料编码
              a1.FNAME,                                 --物料名称
              a1.FSPECIFICATION,                        --物料规格
              a2.FERPCLSID,
              case
                  when a2.FERPCLSID = 1 then N'外购'
                  when a2.FERPCLSID = 2 then N'自制'
                  when a2.FERPCLSID = 3 then N'委外'
                  else N'未定义' end as FErpClass,      --物料属性
              a2.FSUITE,                                --套件
              a2.FBASEUNITID,--物料基本单位内码
              a3.FNUMBER             as FBaseUnitNumber,--物料基本单位
              a2.FWEIGHTUNITID,                         --物料重量单位
              a2.FVOLUMEUNITID,--物料尺寸单位
              a2.FISPURCHASE,--允许采购
              a2.FISINVENTORY,--允许库存
              a2.FISSUBCONTRACT,--允许委外
              a2.FISSALE,--允许销售
              a2.FISPRODUCE,--允许生产
              a2.FISASSET,--允许资产
              a4.FISBATCHMANAGE,                        --是否启用批号管理
              a4.FISKFPERIOD,                           --是否启用保质期管理
              a5.FCHECKINCOMING,--来料检验
              a5.FCHECKPRODUCT,--产品检验,
              a5.FCHECKSTOCK,--库存检验,
              a5.FCHECKRETURN,--退货检验
              a5.FCHECKDELIVERY,--发货检验
              a41.FSTOCKID,                             --默认仓库
              a41.FNAME              as FStockName,     --默认仓库名称
              a.FUSEORGID,                              --使用组织内码
              b.FNUMBER              as FUseOrgNumber,  --使用组织编码
              c.FNAME                as FUseOrgName     --使用组织名称
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
         left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID = a.FUSEORGID
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052


select top 10 *
from T_BD_MATERIAL
where FMASTERID = 118405

select top 10 a.FMASTERID as FMATERIALID, a3.FUSEORGID, a.FUSEORGID
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
         left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID in (a.FUSEORGID, 0)
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052
  and a.FMASTERID = 118405


select top 10 *
from T_BD_UNIT
where FNUMBER = 'zhi'--FUSEORGID = 0
select top 100 *
from xkPdaServer_userInfo_tool
select top 100 *
from xkPdaServer_keeperInfo_tool
select top 100 *
from xkPdaServer_empInfo_tool
select top 100 *
from xkPdaServer_customer_tool
select top 100 *
from xkPdaServer_sale_tz_to_stockout_tool
select top 10 *
from xkPdaServer_sub_ppbom_to_stockout_tool
select *
from xkPdaServer_poorder_to_stockin_tool

select top 10000 a.FSTOCKORGID         as 仓库组织,
                 a.FKEEPERID           as 保管组织,
                 b.FNUMBER             as 仓库编码,
                 c.FNAME               as 仓库名称,
                 isnull(d.FNUMBER, '') as 批号编码,
                 e.FMATERIALID         as 编码内码,
                 e.FNUMBER             as 编码编号,
                 f.FNAME               as 编码名称,
                 g.FNUMBER             as 单位编码,
                 a.FUPDATETIME         as 最后入库时间
from T_STK_INVENTORY a
         left join
     T_BD_LOTMASTER d on a.FLOT = d.FLOTID,
     T_BD_STOCK b,
     T_BD_STOCK_L c,
     T_BD_MATERIAL e,
     T_BD_MATERIAL_L f,
     T_BD_UNIT g
where a.FSTOCKID = b.FSTOCKID
  and b.FSTOCKID = c.FSTOCKID
  and c.FLOCALEID = 2052
  and a.FMATERIALID = e.FMATERIALID
  and e.FMATERIALID = f.FMATERIALID
  and f.FLOCALEID = 2052
  and a.FSTOCKSTATUSID = '10000'
  and a.FBASEUNITID = g.FMASTERID
  and a.FBASEQTY <> 0


select top 10 *
from T_STK_INVENTORY


select top 10 *
from xkPdaServer_gxplan_to_gxreport_tool

select top 10 *
from T_SFC_OPERPLANNING


select top 10 a.FDEPARTMENTID as 加工车间内码, b.fnumber, *
from T_SFC_OPERPLANNINGDETAIL a, xkPdaServer_depart_tool b where a.FDEPARTMENTID = b.FMASTERID

select a.FMASTERID, a.FNUMBER, b.FNAME, c.FNUMBER as FUseOrgNumber
from T_BD_DEPARTMENT a,
     T_BD_DEPARTMENT_L b,
     T_ORG_organizations c
where a.FMASTERID = b.FDEPTID
  and b.FLOCALEID = 2052
  and a.FDOCUMENTSTATUS = 'C'
and a.FUSEORGID = c.FORGID
select * from    xkPdaServer_orgInfo_tool

select * from xkPdaServer_gxplan_to_gxreport_tool
